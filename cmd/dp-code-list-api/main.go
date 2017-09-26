package main

import (
	"os"

	"context"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/api"
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-code-list-api/mongostore"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/go-ns/mongo"
	"github.com/gorilla/mux"
	"os/signal"
	"syscall"
	"errors"
	"net/http"
)

func main() {
	log.Namespace = "dp-code-list-api"

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Get()
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	mongoDatastore, err := mongostore.CreateMongoDataStore(cfg.MongoDBURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	router := mux.NewRouter()
	httpErrChannel := make(chan error)
	_ = api.CreateCodeListAPI(router, mongoDatastore)
	httpServer := server.New(cfg.BindAddr, router)
	httpServer.HandleOSSignals = false

	shutdown := func(httpShutdown bool) {
		log.Info(fmt.Sprintf("shutdown with timeout: %d", cfg.GracefulShutdownTimeout), nil)
		ctx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

		if httpShutdown {
			err = httpServer.Shutdown(ctx)
			if err != nil {
				log.Error(err, nil)
			}
		}

		///mongo.Close() may use all remaining time in the context - do this last!
		if err = mongo.Close(ctx, mongoDatastore.Session); err != nil {
			log.Error(err, nil)
		}

		log.Info("shutdown complete", nil)

		cancel()
		os.Exit(1)
	}

	go func() {
		log.Info("code list API listening .....", log.Data{"BIND_ADDR": cfg.BindAddr})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error(err, nil)
			httpErrChannel <- err
		} else {
			httpErrChannel <- errors.New("http completed shutdown")
		}
	}()

	for {
		select {
		case err := <-httpErrChannel:
			log.ErrorC("http error received", err, nil)
			shutdown(false)
		case <-signals:
			log.Debug("os signal received", nil)
			shutdown(true)
		}
	}
}
