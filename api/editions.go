package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/models"
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

	dbEditions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		data["err"] = err
		log.Event(ctx, "failed to get editions", log.INFO, data)
		handleError(ctx, w, err, data)
		return
	}
	editions := models.NewEditions(dbEditions)

	for i, item := range editions.Items {
		if err := item.UpdateLinks(id, c.apiURL); err != nil {
			log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getEditions endpoint: links could not be created")))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
		editions.Items[i] = item
	}

	count := len(editions.Items)
	editions.Count = count
	editions.Limit = count
	editions.TotalCount = count

	b, err := json.Marshal(editions)
	if err != nil {
		data["err"] = err
		log.Event(ctx, "failed to marshal editions", log.INFO, data)
		handleError(ctx, w, err, data)
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
		data["err"] = err
		log.Event(ctx, "failed to get edition", log.INFO, data)
		handleError(ctx, w, err, data)
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
		data["err"] = err
		log.Event(ctx, "failed to marshal editions", log.INFO, data)
		handleError(ctx, w, err, data)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getEdition endpoint: failed to write bytes to response")), data)
		return
	}
	log.Event(r.Context(), "retrieved codelist edition", log.INFO, log.Data{"code_list_id": id, edition: edition})
}
