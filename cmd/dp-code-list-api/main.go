package main

import (
	"os"

	"context"
	"errors"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-code-list-api/api"
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-graph/graph"
	"github.com/ONSdigital/go-ns/healthcheck"
	"github.com/ONSdigital/go-ns/server"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
)

var (
	// BuildTime represents the time in which the service was built
	BuildTime string
	// GitCommit represents the commit (SHA-1) hash of the service that is running
	GitCommit string
	// Version represents the version of the service that is running
	Version string
)

func main() {
	log.Namespace = "dp-code-list-api"
	ctx := context.Background()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	cfg, err := config.Get()
	if err != nil {
		log.Event(ctx, "error getting config", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	datastore, err := graph.NewCodeListStore(context.Background())
	if err != nil {
		log.Event(ctx, "error creating codelist store", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	router := mux.NewRouter()
	router.Path("/healthcheck").HandlerFunc(healthcheck.Do)
	httpErrChannel := make(chan error)
	api.CreateCodeListAPI(router, datastore, cfg.CodeListAPIURL, cfg.DatasetAPIURL)
	httpServer := server.New(cfg.BindAddr, router)
	httpServer.HandleOSSignals = false

	healthTicker := healthcheck.NewTicker(
		cfg.HealthCheckInterval,
		cfg.HealthCheckRecovery,
		datastore,
	)

	shutdown := func(httpShutdown bool) {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

		log.Event(shutdownCtx, "shutdown with timeout", log.INFO, log.Data{"timeout": cfg.GracefulShutdownTimeout})

		if httpShutdown {
			err = httpServer.Shutdown(shutdownCtx)
			if err != nil {
				log.Event(shutdownCtx, "http server shutdown error", log.ERROR, log.Error(err))
			}
		}

		healthTicker.Close()

		if err = datastore.Close(shutdownCtx); err != nil {
			log.Event(shutdownCtx, "datastore close error", log.ERROR, log.Error(err))
		}

		log.Event(shutdownCtx, "shutdown complete", log.INFO)

		cancel()
		os.Exit(1)
	}

	go func() {
		log.Event(ctx, "code list api starting.....", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Event(ctx, "error starting http server", log.ERROR, log.Error(err))
			httpErrChannel <- err
		} else {
			httpErrChannel <- errors.New("http completed shutdown")
		}
	}()

	for {
		select {
		case err := <-httpErrChannel:
			log.Event(ctx, "http error received", log.ERROR, log.Error(err))
			shutdown(false)
		case <-signals:
			log.Event(ctx, "os signal received", log.INFO)
			shutdown(true)
		}
	}
}
