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
	dphttp "github.com/ONSdigital/dp-net/v2/http"
	dpotelgo "github.com/ONSdigital/dp-otel-go"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
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
		log.Error(ctx, "error getting config", err)
		os.Exit(1)
	}

	// Create CodeList Store
	datastore, err := graph.NewCodeListStore(ctx)
	if err != nil {
		log.Error(ctx, "error creating codelist store", err)
		os.Exit(1)
	}

	graphErrorConsumer := graph.NewLoggingErrorConsumer(ctx, datastore.Errors)

	// Create healthcheck object with versionInfo
	versionInfo, err := healthcheck.NewVersionInfo(BuildTime, GitCommit, Version)
	if err != nil {
		log.Error(ctx, "failed to create versionInfo for healthcheck", err)
		os.Exit(1)
	}
	hc := healthcheck.New(versionInfo, cfg.HealthCheckCriticalTimeout, cfg.HealthCheckInterval)

	// Register checkers
	if err := registerCheckers(ctx, &hc, datastore); err != nil {
		os.Exit(1)
	}

	// Set up OpenTelemetry
	otelConfig := dpotelgo.Config{
		OtelServiceName:          cfg.OTServiceName,
		OtelExporterOtlpEndpoint: cfg.OTExporterOTLPEndpoint,
		OtelBatchTimeout:         cfg.OTBatchTimeout,
	}

	otelShutdown, err := dpotelgo.SetupOTelSDK(ctx, otelConfig)

	if err != nil {
		log.Error(ctx, "error setting up OpenTelemetry - hint: ensure OTEL_EXPORTER_OTLP_ENDPOINT is set", err)
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	// Create HTTP Server with health endpoint and CodeList API
	router := mux.NewRouter()
	router.Use(otelmux.Middleware(cfg.OTServiceName))
	router.Path("/health").HandlerFunc(hc.Handler)

	api.CreateCodeListAPI(router, datastore, cfg.CodeListAPIURL, cfg.DatasetAPIURL, cfg.DefaultOffset, cfg.DefaultLimit, cfg.DefaultMaxLimit)
	otelHandler := otelhttp.NewHandler(router, "/")
	httpServer := dphttp.NewServer(cfg.BindAddr, otelHandler)
	httpServer.HandleOSSignals = false

	// Start healthcheck ticker
	hc.Start(ctx)

	// Start HTTP Server
	go func() {
		log.Info(ctx, "code list api starting.....", log.Data{"bind_addr": cfg.BindAddr})
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			hc.Stop()
			log.Error(ctx, "error starting http server", err)
		}
	}()

	// wait until we receive a signal
	<-signals
	log.Info(ctx, "os signal received")

	// Shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), cfg.GracefulShutdownTimeout)

	// Graceful shutdown function
	shutdown := func() {
		anyError := false
		log.Info(shutdownCtx, "shutdown with timeout", log.Data{"timeout": cfg.GracefulShutdownTimeout})

		// StopHTTP Server
		err = httpServer.Shutdown(shutdownCtx)
		if err != nil {
			anyError = true
			log.Error(shutdownCtx, "http server shutdown error", err)
		} else {
			log.Info(shutdownCtx, "http server successful shutdown")
		}

		// Stop healthcheck
		hc.Stop()
		log.Info(shutdownCtx, "healthcheck stopped")

		// Close data store
		if err = datastore.Close(shutdownCtx); err != nil {
			anyError = true
			log.Error(shutdownCtx, "datastore close error", err)
		} else {
			log.Info(shutdownCtx, "datastore successfully closed")
		}

		if err = graphErrorConsumer.Close(shutdownCtx); err != nil {
			anyError = true
			log.Error(shutdownCtx, "graph error consumer close error", err)
		} else {
			log.Info(shutdownCtx, "graph error consumer successfully closed")
		}

		// If any error happened during shutdown, log it and exit with err code
		if anyError {
			log.Warn(ctx, "graceful shutdown had errors")
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
		log.Error(shutdownCtx, "shutdown timeout", shutdownCtx.Err())
		os.Exit(1)
	}
	log.Info(ctx, "graceful shutdown was successful")
	os.Exit(0)
}

// RegisterCheckers adds the checkers for the provided clients to the healthcheck object.
func registerCheckers(ctx context.Context, hc *healthcheck.HealthCheck, db *graph.DB) (err error) {
	hasErrors := false

	if err = hc.AddCheck("Graph DB", db.Checker); err != nil {
		hasErrors = true
		log.Error(ctx, "error adding check for graph db", err)
	}

	if hasErrors {
		return errors.New("error registering checkers for health check")
	}
	return nil
}
