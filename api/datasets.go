package api

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
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
	offsetParameter := r.URL.Query().Get("offset")
	limitParameter := r.URL.Query().Get("limit")
	offset := c.defaultOffset
	limit := c.defaultLimit
	var err error

	log.Event(ctx, "getCodeDatasets endpoint: attempting to find datasets related to code", log.INFO, logData)

	if offsetParameter != "" {
		logData["offset"] = offsetParameter
		offset, err = ValidatePositiveInt(offsetParameter)
		if err != nil {
			log.Event(ctx, "invalid query parameter: offset", log.ERROR, log.Error(err), logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limitParameter != "" {
		logData["limit"] = limitParameter
		limit, err = ValidatePositiveInt(limitParameter)
		if err != nil {
			log.Event(ctx, "invalid query parameter: limit", log.ERROR, log.Error(err), logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limit > c.maxLimit {
		logData["max_limit"] = c.maxLimit
		err = errors.New("limit is greater than the maximum allowed")
		log.Event(ctx, "invalid query parameter: limit, maximum limit reached", log.ERROR, log.Error(err), logData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbDatasets, err := c.store.GetCodeDatasets(ctx, codeListID, edition, code)
	if err != nil {
		handleError(ctx, "failed to get datasets list", logData, err, w)
		return
	}

	totalCount := len(dbDatasets.Items)

	sort.Slice(dbDatasets.Items, func(i, j int) bool {
		return dbDatasets.Items[i].ID < dbDatasets.Items[j].ID
	})

	slicedResults := datasetsSlice(dbDatasets.Items, offset, limit)

	datasets := models.NewDatasets(slicedResults)

	if err := datasets.UpdateLinks(c.datasetAPIURL, codeListID); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodeDatasets endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

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

func datasetsSlice(full []dbmodels.Dataset, offset, limit int) (sliced []dbmodels.Dataset) {
	end := offset + limit
	if end > len(full) {
		end = len(full)
	}

	if offset > len(full) || limit == 0 {
		return []dbmodels.Dataset{}
	}
	return full[offset:end]
}
