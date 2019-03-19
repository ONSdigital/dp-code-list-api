package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getEditions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	data := log.Data{"codelist_id": id}

	log.InfoCtx(ctx, "getEditions endpoint: attempting to find editions", data)

	editions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		data["err"] = err
		log.InfoCtx(ctx, "failed to get editions", data)
		handleError(ctx, w, err, data)
		return
	}

	for _, item := range editions.Items {
		if err := item.UpdateLinks(id, c.apiURL); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "getEditions endpoint: links could not be created"), nil)
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
	}

	count := len(editions.Items)
	editions.Count = count
	editions.Limit = count
	editions.TotalCount = count

	b, err := json.Marshal(editions)
	if err != nil {
		data["err"] = err
		log.InfoCtx(ctx, "failed to marshal editions", data)
		handleError(ctx, w, err, data)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getEditions endpoint: failed to write bytes to response"), data)
		return
	}
	log.InfoCtx(r.Context(), "retrieved codelist editions", log.Data{"code_list_id": id})
}

func (c *CodeListAPI) getEdition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.InfoCtx(ctx, "getEdition endpoint: attempting to find edition", data)

	editionModel, err := c.store.GetEdition(r.Context(), id, edition)
	if err != nil {
		data["err"] = err
		log.InfoCtx(ctx, "failed to get edition", data)
		handleError(ctx, w, err, data)
		return
	}

	if err := editionModel.UpdateLinks(id, c.apiURL); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getEdition endpoint: links could not be created"), nil)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(editionModel)
	if err != nil {
		data["err"] = err
		log.InfoCtx(ctx, "failed to marshal editions", data)
		handleError(ctx, w, err, data)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getEdition endpoint: failed to write bytes to response"), data)
		return
	}
	log.InfoCtx(r.Context(), "retrieved codelist edition", log.Data{"code_list_id": id, edition: edition})
}
