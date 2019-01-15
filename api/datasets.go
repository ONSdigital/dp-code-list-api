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
	codeListID := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	logData := log.Data{"code_list_id": codeListID, "edition": edition, "code": code}

	log.InfoCtx(ctx, "getCodeDatasets endpoint: attempting to find datasets related to code", logData)

	datasets, err := c.store.GetCodeDatasets(ctx, codeListID, edition, code)
	if err != nil {
		handleError(ctx, w, err, logData)
		return
	}

	if err := datasets.UpdateLinks(c.apiURL, c.datasetAPIURL, codeListID, edition, code); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeDatasets endpoint: links could not be created"), nil)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	count := len(datasets.Items)
	datasets.Count = count
	datasets.Limit = count
	datasets.TotalCount = count

	b, err := json.Marshal(datasets)
	if err != nil {
		handleError(ctx, w, err, logData)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		handleError(ctx, w, err, logData)
		return
	}

	log.InfoCtx(ctx, "getCodeDatasets endpoint: request successful", logData)
}
