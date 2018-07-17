package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getCodeDatasets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	logData := log.Data{"code_list_id": id, "edition": edition, "code": code}

	log.InfoCtx(ctx, "getCodeDatasets endpoint: attempting to find datasets related to code", logData)

	datasets, err := c.store.GetCodeDatasets(ctx, id, edition, code)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeDatasets endpoint: could not get code datasets from store"), logData)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(datasets)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeDatasets endpoint: failed to marshal datasets to json bytes"), logData)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeDatasets endpoint: failed to write bytes to response"), logData)
		return
	}

	log.InfoCtx(ctx, "getCodeDatasets endpoint: request successful", logData)
}
