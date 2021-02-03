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

func (c *CodeListAPI) getEditions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	data := log.Data{"codelist_id": id}

	log.Event(ctx, "getEditions endpoint: attempting to find editions", log.INFO, data)

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

	dbEditions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		handleError(ctx, "failed to get editions", data, err, w)
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
			log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getEditions endpoint: links could not be created")))
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
		handleError(ctx, "failed to marshal editions", data, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getEditions endpoint: failed to write bytes to response")), data)
		return
	}
	log.Event(r.Context(), "retrieved codelist editions", log.INFO, log.Data{"code_list_id": id})
}

func (c *CodeListAPI) getEdition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.Event(ctx, "getEdition endpoint: attempting to find edition", log.INFO, data)

	dbEditionModel, err := c.store.GetEdition(r.Context(), id, edition)
	if err != nil {
		handleError(ctx, "failed to get edition", data, err, w)
		return
	}
	editionModel := models.NewEdition(dbEditionModel)

	if err := editionModel.UpdateLinks(id, c.apiURL); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getEdition endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(editionModel)
	if err != nil {
		handleError(ctx, "failed to marshal editions", data, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getEdition endpoint: failed to write bytes to response")), data)
		return
	}
	log.Event(r.Context(), "retrieved codelist edition", log.INFO, log.Data{"code_list_id": id, edition: edition})
}

func editionsSlice(full []dbmodels.Edition, offset, limit int) (sliced []dbmodels.Edition) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []dbmodels.Edition{}
	}
	return full[offset:end]
}
