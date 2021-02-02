package api

import (
	"encoding/json"
	"net/http"
	"sort"

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

	// get limit from query parameters, or default value
	limit, err := GetPositiveIntQueryParameter(r.URL.Query(), "limit", c.defaultLimit)
	if err != nil {
		log.Event(ctx, "failed to obtain limit from request query parameters", log.ERROR)
		handleError(ctx, "failed to obtain limit", log.Data{"limit": limit}, err, w)
		return
	}

	// get offset from query parameters, or default value
	offset, err := GetPositiveIntQueryParameter(r.URL.Query(), "offset", c.defaultOffset)
	if err != nil {
		log.Event(ctx, "failed to obtain offset from request query parameters", log.ERROR)
		handleError(ctx, "failed to obtain offset", log.Data{"offset": offset}, err, w)
		return
	}

	dbDatasets, err := c.store.GetCodeDatasets(ctx, codeListID, edition, code)
	if err != nil {
		handleError(ctx, "failed to get datasets list", logData, err, w)
		return
	}

	totalCount := len(dbDatasets.Items)

	datasets := models.NewDatasets(dbDatasets)

	if err := datasets.UpdateLinks(c.datasetAPIURL, codeListID); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodeDatasets endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	sort.Slice(datasets.Items, func(i, j int) bool {
		return datasets.Items[i].ID < datasets.Items[j].ID
	})

	slicedResults := datasetsSlice(datasets.Items, offset, limit)

	count := len(slicedResults)
	datasets.Count = count
	datasets.Offset = offset
	datasets.Limit = limit
	datasets.TotalCount = totalCount

	b, err := json.Marshal(datasets)
	if err != nil {
		handleError(ctx, "failed to marshal datasets list", logData, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		handleError(ctx, "failed to write body", logData, err, w)
		return
	}

	log.Event(ctx, "getCodeDatasets endpoint: request successful", log.INFO, logData)
}

func datasetsSlice(full []models.Dataset, offset, limit int) (sliced []models.Dataset) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []models.Dataset{}
	}
	return full[offset:end]
}
