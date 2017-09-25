package main

import (
	"os"

	"github.com/ONSdigital/dp-code-list-api/api"
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-code-list-api/mongo"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "dp-code-list-api"
	cfg, err := config.Get()
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	mongoDatastore, err := mongo.CreateMongoDataStore(cfg.MongoDBURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	router := mux.NewRouter()
	_ = api.CreateCodeListAPI(router, mongoDatastore)
	log.Info("code List API listening .....", log.Data{"BIND_ADDR": cfg.BindAddr})
	srv := server.New(cfg.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}
