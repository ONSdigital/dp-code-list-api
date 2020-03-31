package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/log.go/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getCodeDatasets(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	codeListID := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	logData := log.Data{"code_list_id": codeListID, "edition": edition, "code": code}

	log.Event(ctx, "getCodeDatasets endpoint: attempting to find datasets related to code", log.INFO, logData)

	dbDatasets, err := c.store.GetCodeDatasets(ctx, codeListID, edition, code)
	if err != nil {
		logData["err"] = err
		log.Event(ctx, "failed to get datasets list", log.INFO, logData)
		handleError(ctx, w, err, logData)
		return
	}
	datasets := models.NewDatasets(dbDatasets)

	if err := datasets.UpdateLinks(c.datasetAPIURL, codeListID); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodeDatasets endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	count := len(datasets.Items)
	datasets.Count = count
	datasets.Limit = count
	datasets.TotalCount = count

	b, err := json.Marshal(datasets)
	if err != nil {
		logData["err"] = err
		log.Event(ctx, "failed to marshal datasets list", log.INFO, logData)
		handleError(ctx, w, err, logData)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		logData["err"] = err
		log.Event(ctx, "failed to write body", log.INFO, logData)
		handleError(ctx, w, err, logData)
		return
	}

	log.Event(ctx, "getCodeDatasets endpoint: request successful", log.INFO, logData)
}
