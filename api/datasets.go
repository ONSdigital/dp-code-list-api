package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
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
		logData["err"] = err
		log.InfoCtx(ctx, "failed to get datasets list", logData)
		handleError(ctx, w, err, logData)
		return
	}

	b, err := json.Marshal(datasets)
	if err != nil {
		logData["err"] = err
		log.InfoCtx(ctx, "failed to marshal datasets list", logData)
		handleError(ctx, w, err, logData)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		logData["err"] = err
		log.InfoCtx(ctx, "failed to write body", logData)
		handleError(ctx, w, err, logData)
		return
	}

	log.InfoCtx(ctx, "getCodeDatasets endpoint: request successful", logData)
}
