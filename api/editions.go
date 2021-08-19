package api

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getEditions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	logData := log.Data{"codelist_id": id}
	offsetParameter := r.URL.Query().Get("offset")
	limitParameter := r.URL.Query().Get("limit")
	offset := c.defaultOffset
	limit := c.defaultLimit
	var err error

	log.Info(ctx, "getEditions endpoint: attempting to find editions", logData)

	if offsetParameter != "" {
		logData["offset"] = offsetParameter
		offset, err = ValidatePositiveInt(offsetParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: offset", err, logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limitParameter != "" {
		logData["limit"] = limitParameter
		limit, err = ValidatePositiveInt(limitParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: limit", err, logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limit > c.maxLimit {
		logData["max_limit"] = c.maxLimit
		err = errors.New("limit is greater than the maximum allowed")
		log.Error(ctx, "invalid query parameter: limit, maximum limit reached", err, logData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbEditions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		handleError(ctx, "failed to get editions", logData, err, w)
		return
	}

	totalCount := len(dbEditions.Items)

	sort.Slice(dbEditions.Items, func(i, j int) bool {
		return dbEditions.Items[i].ID < dbEditions.Items[j].ID
	})

	slicedResults := editionsSlice(dbEditions.Items, offset, limit)

	editions := models.NewEditions(slicedResults)

	for i, item := range editions.Items {
		if err := item.UpdateLinks(id, c.apiURL); err != nil {
			log.Error(ctx, "error updating links", errors.WithMessage(err, "getEditions endpoint: links could not be created"))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
		editions.Items[i] = item
	}

	count := len(slicedResults)
	editions.Count = count
	editions.Offset = offset
	editions.Limit = limit
	editions.TotalCount = totalCount

	b, err := json.Marshal(editions)
	if err != nil {
		handleError(ctx, "failed to marshal editions", logData, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Error(ctx, "error writting body", errors.WithMessage(err, "getEditions endpoint: failed to write bytes to response"), logData)
		return
	}
	log.Info(r.Context(), "retrieved codelist editions", log.Data{"code_list_id": id})
}

func (c *CodeListAPI) getEdition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.Info(ctx, "getEdition endpoint: attempting to find edition", data)

	dbEditionModel, err := c.store.GetEdition(r.Context(), id, edition)
	if err != nil {
		handleError(ctx, "failed to get edition", data, err, w)
		return
	}
	editionModel := models.NewEdition(dbEditionModel)

	if err := editionModel.UpdateLinks(id, c.apiURL); err != nil {
		log.Error(ctx, "error updating links", errors.WithMessage(err, "getEdition endpoint: links could not be created"))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(editionModel)
	if err != nil {
		handleError(ctx, "failed to marshal editions", data, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Error(ctx, "error writting body", errors.WithMessage(err, "getEdition endpoint: failed to write bytes to response"), data)
		return
	}
	log.Info(r.Context(), "retrieved codelist edition", log.Data{"code_list_id": id, edition: edition})
}

func editionsSlice(full []dbmodels.Edition, offset, limit int) (sliced []dbmodels.Edition) {
	end := offset + limit
	if end > len(full) {
		end = len(full)
	}

	if offset > len(full) || limit == 0 {
		return []dbmodels.Edition{}
	}
	return full[offset:end]
}
