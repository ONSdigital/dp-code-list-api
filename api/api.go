package api

import (
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
)

// CodeListAPI holds all endpoints which are used to access the code list resources
type CodeListAPI struct {
	router *mux.Router
	store  datastore.DataStore
}

// CreateCodeListAPI returns a constructed code list api
func CreateCodeListAPI(route *mux.Router, store datastore.DataStore) *CodeListAPI {
	api := CodeListAPI{router: route, store: store}
	api.router.HandleFunc("/code-lists", api.getCodeLists).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}", api.getCodeList).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions", api.getEditions).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}", api.getEdition).Methods("GET")
	return &api
}

func handleError(w http.ResponseWriter, err error) {
	log.Error(err, nil)
	if datastore.NOT_FOUND == err {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func writeBody(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(bytes); err != nil {
		log.Error(err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
