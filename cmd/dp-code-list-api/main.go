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
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
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

	// Get Config
	cfg, err := config.Get()
	if err != nil {
		log.Event(ctx, "error getting config", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	// Create CodeList Store
	datastore, err := graph.NewCodeListStore(context.Background())
	if err != nil {
		log.Event(ctx, "error creating codelist store", log.ERROR, log.Error(err))
		os.Exit(1)
	}

	// Create healthcheck object with versionInfo
	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	if err != nil {
		log.Event(ctx, "failed to create versionInfo for healthcheck", log.FATAL, log.Error(err))
		os.Exit(1)
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

	// Register checkers
	if err := registerCheckers(ctx, &hc); err != nil {
		os.Exit(1)
	}

	// Create HTTP Server with health endpoint and CodeList API
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(hc.Handler)
	httpErrChannel := make(chan error)
	api.CreateCodeListAPI(router, datastore, cfg.CodeListAPIURL, cfg.DatasetAPIURL)
	httpServer := server.New(cfg.BindAddr, router)
	httpServer.HandleOSSignals = false

	// Graceful shutdown function
	shutdown := func() {

		shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)
		anyError := false
		log.Event(shutdownCtx, "shutdown with timeout", log.INFO, log.Data{"timeout": cfg.GracefulShutdownTimeout})

		// StopHTTP Server
		err = httpServer.Shutdown(shutdownCtx)
		if err != nil {
			anyError = true
			log.Event(shutdownCtx, "http server shutdown error", log.ERROR, log.Error(err))
		} else {
			log.Event(shutdownCtx, "http server successful shutdown", log.INFO)
		}

		// Strop healthcheck
		hc.Stop()
		log.Event(shutdownCtx, "healthcheck stopped", log.INFO)

		// Close data store
		if err = datastore.Close(shutdownCtx); err != nil {
			anyError = true
			log.Event(shutdownCtx, "datastore close error", log.ERROR, log.Error(err))
		} else {
			log.Event(shutdownCtx, "datastore successfully closed", log.INFO)
		}

		// cancel the timer in the shutdown context.
		cancel()

		// if any error happened during shutdown, log it and exit with err code
		if anyError {
			log.Event(ctx, "graceful shutdown had errors", log.WARN)
			os.Exit(1)
		}

		// if all dependencies shutted down successfully, log it and exit with success code
		log.Event(ctx, "graceful shutdown was successful", log.INFO)
		os.Exit(0)
	}

	// Start healthcheck ticker
	hc.Start(ctx)

	// Start HTTP Server
	go func() {
		log.Event(ctx, "code list api starting.....", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hc.Stop()
			log.Event(ctx, "error starting http server", log.ERROR, log.Error(err))
			httpErrChannel <- err
		} else {
			httpErrChannel <- errors.New("http completed shutdown")
		}
	}()

	// log errors
	go func() {
		for {
			select {
			case err := <-httpErrChannel:
				log.Event(ctx, "http error received", log.ERROR, log.Error(err))
			}
		}
	}()

	// wait until we receive a signal
	<-signals
	log.Event(ctx, "os signal received", log.INFO)
	shutdown()
}

// RegisterCheckers adds the checkers for the provided clients to the healthcheck object.
func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck) (err error) {

	hasErrors := false

	if hasErrors {
		return errors.New("Error(s) registering checkers for healthcheck")
	}
	return nil
}
