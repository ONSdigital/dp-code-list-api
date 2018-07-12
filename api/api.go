package api

import (
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"context"
)

const (
	internalServerErr = "internal server error"
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

// CodeListAPI holds all endpoints which are used to access the code list resources
type CodeListAPI struct {
	router    *mux.Router
	store     datastore.DataStore
	writeBody func(w http.ResponseWriter, bytes []byte) error
}

// CreateCodeListAPI returns a constructed code list api
func CreateCodeListAPI(route *mux.Router, store datastore.DataStore) *CodeListAPI {
	api := CodeListAPI{
		router: route,
		store:  store,
		writeBody: func(w http.ResponseWriter, bytes []byte) error {
			w.Header().Set(contentTypeHeader, contentTypeJSON)
			if _, err := w.Write(bytes); err != nil {
				http.Error(w, internalServerErr, http.StatusInternalServerError)
				return err
			}
			return nil
		},
	}

	api.router.HandleFunc("/code-lists", api.getCodeLists).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}", api.getCodeList).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions", api.getEditions).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}", api.getEdition).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}/codes", api.getCodes).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}/codes/{code}", api.getCode).Methods("GET")
	return &api
}

func handleError(ctx context.Context, w http.ResponseWriter, err error, data log.Data) {
	log.ErrorCtx(ctx, err, data)
	if datastore.NOT_FOUND == err {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
}

func writeBody(w http.ResponseWriter, bytes []byte) error {
	w.Header().Set(contentTypeHeader, contentTypeJSON)
	if _, err := w.Write(bytes); err != nil {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return err
	}
	return nil
}
