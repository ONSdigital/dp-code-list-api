package main

import (
	"net/http"
	"os"

	"context"
	"errors"
	"os/signal"
	"syscall"

	"github.com/ONSdigital/dp-code-list-api/api"
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-graph/v2/graph"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
	dphttp "github.com/ONSdigital/dp-net/http"
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
		log.Event(ctx, "error getting config", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	// Create CodeList Store
	datastore, err := graph.NewCodeListStore(ctx)
	if err != nil {
		log.Event(ctx, "error creating codelist store", log.FATAL, log.Error(err))
		os.Exit(1)
	}

	graphErrorConsumer := graph.NewLoggingErrorConsumer(ctx, datastore.Errors)

	// Create healthcheck object with versionInfo
	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	if err != nil {
		log.Event(ctx, "failed to create versionInfo for healthcheck", log.FATAL, log.Error(err))
		os.Exit(1)
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

	// Register checkers
	if err := registerCheckers(ctx, &hc, datastore); err != nil {
		os.Exit(1)
	}

	// Create HTTP Server with health endpoint and CodeList API
	router := mux.NewRouter()
	router.Path("/health").HandlerFunc(hc.Handler)

	api.CreateCodeListAPI(router, datastore, cfg.CodeListAPIURL, cfg.DatasetAPIURL, cfg.DefaultOffset, cfg.DefaultLimit)
	httpServer := dphttp.NewServer(cfg.BindAddr, router)
	httpServer.HandleOSSignals = false

	// Start healthcheck ticker
	hc.Start(ctx)

	// Start HTTP Server
	go func() {
		log.Event(ctx, "code list api starting.....", log.INFO, log.Data{"bind_addr": cfg.BindAddr})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hc.Stop()
			log.Event(ctx, "error starting http server", log.ERROR, log.Error(err))
		}
	}()

	// wait until we receive a signal
	<-signals
	log.Event(ctx, "os signal received", log.INFO)

	// Shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

	// Graceful shutdown function
	shutdown := func() {

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

		// Stop healthcheck
		hc.Stop()
		log.Event(shutdownCtx, "healthcheck stopped", log.INFO)

		// Close data store
		if err = datastore.Close(shutdownCtx); err != nil {
			anyError = true
			log.Event(shutdownCtx, "datastore close error", log.ERROR, log.Error(err))
		} else {
			log.Event(shutdownCtx, "datastore successfully closed", log.INFO)
		}

		if err = graphErrorConsumer.Close(shutdownCtx); err != nil {
			anyError = true
			log.Event(shutdownCtx, "graph error consumer close error", log.ERROR, log.Error(err))
		} else {
			log.Event(shutdownCtx, "graph error consumer successfully closed", log.INFO)
		}

		// If any error happened during shutdown, log it and exit with err code
		if anyError {
			log.Event(ctx, "graceful shutdown had errors", log.WARN)
			os.Exit(1)
		}

		// cancel the timer in the shutdown context once everything is shutted down successfully.
		cancel()
	}

	// Perform shutdown in parallel go-routine
	go shutdown()

	// Wait for Shutdown timeout or success (via cancel)
	<-shutdownCtx.Done()
	if shutdownCtx.Err() == context.DeadlineExceeded {
		log.Event(shutdownCtx, "shutdown timeout", log.ERROR, log.Error(shutdownCtx.Err()))
		os.Exit(1)
	}
	log.Event(ctx, "graceful shutdown was successful", log.INFO)
	os.Exit(0)
}

// RegisterCheckers adds the checkers for the provided clients to the healthcheck object.
func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck, db *graph.DB) (err error) {

	hasErrors := false

	if err = hc.AddCheck("Graph DB", db.Checker); err != nil {
		hasErrors = true
		log.Event(ctx, "error adding check for graph db", log.ERROR, log.Error(err))
	}

	if hasErrors {
		return errors.New("error registering checkers for health check")
	}
	return nil
}
