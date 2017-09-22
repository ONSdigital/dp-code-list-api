package api

import (
	"encoding/json"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"net/http"
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
	api.router.HandleFunc("/code-lists/{id}/codes", api.getCodes).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/codes/{code}", api.getCode).Methods("GET")
	return &api
}

func (c *CodeListAPI) getCodeLists(w http.ResponseWriter, r *http.Request) {

	codeLists, err := c.store.GetCodeLists()
	if err != nil {
		log.Error(err, nil)
		return
	}

	bytes, err := json.Marshal(codeLists)
	if err != nil {
		internalError(w, err)
		return
	}

	writeBody(w, bytes)
	log.Debug("get all instances", nil)
}

func (c *CodeListAPI) getCodeList(w http.ResponseWriter, r *http.Request) {

}

func (c *CodeListAPI) getCodes(w http.ResponseWriter, r *http.Request) {

}

func (c *CodeListAPI) getCode(w http.ResponseWriter, r *http.Request) {

}

func internalError(w http.ResponseWriter, err error) {
	log.Error(err, nil)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func writeBody(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(bytes); err != nil {
		log.Error(err, nil)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
