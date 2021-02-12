package api

import (
	"net/http"
	"strconv"

	"context"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

const (
	internalServerErr = "internal server error"
	contentTypeHeader = "Content-Type"
	contentTypeJSON   = "application/json"
)

// CodeListAPI holds all endpoints which are used to access the code list resources
type CodeListAPI struct {
	router        *mux.Router
	store         datastore.DataStore
	writeBody     func(w http.ResponseWriter, bytes []byte) error
	apiURL        string
	datasetAPIURL string
	defaultOffset int
	defaultLimit  int
	maxLimit      int
}

// CreateCodeListAPI returns a constructed code list api
func CreateCodeListAPI(route *mux.Router, store datastore.DataStore, apiURL, datasetAPIURL string, defaultOffset, defaultLimit, maxLimit int) *CodeListAPI {
	api := CodeListAPI{
		router:        route,
		store:         store,
		apiURL:        apiURL,
		datasetAPIURL: datasetAPIURL,
		writeBody: func(w http.ResponseWriter, bytes []byte) error {
			w.Header().Set(contentTypeHeader, contentTypeJSON)
			if _, err := w.Write(bytes); err != nil {
				http.Error(w, internalServerErr, http.StatusInternalServerError)
				return err
			}
			return nil
		},
		defaultOffset: defaultOffset,
		defaultLimit:  defaultLimit,
		maxLimit:      maxLimit,
	}

	api.router.HandleFunc("/code-lists", api.getCodeLists).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}", api.getCodeList).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions", api.getEditions).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}", api.getEdition).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}/codes", api.getCodes).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}/codes/{code}", api.getCode).Methods("GET")
	api.router.HandleFunc("/code-lists/{id}/editions/{edition}/codes/{code}/datasets", api.getCodeDatasets).Methods("GET")
	return &api
}

func handleError(ctx context.Context, logMsg string, logData log.Data, err error, w http.ResponseWriter) {
	log.Event(ctx, logMsg, log.ERROR, log.Error(err), logData)
	if err == driver.ErrNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
	}
}

// ValidatePositiveInt obtains the positive int value of query var defined by the provided varKey
func ValidatePositiveInt(parameter string) (val int, err error) {
	err = errors.New("value is not a positive integer")
	val, err = strconv.Atoi(parameter)
	if err != nil {
		return -1, err
	}
	if val < 0 {
		return -1, err
	}
	return val, nil
}
