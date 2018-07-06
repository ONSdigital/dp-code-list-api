package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
)

func (c *CodeListAPI) getEditions(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	data := log.Data{"codelist_id": id}

	editions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		handleError(ctx, w, err, data)
		return
	}

	b, err := json.Marshal(editions)
	if err != nil {
		handleError(ctx, w, err, data)
		return
	}

	writeBody(ctx, w, b)
	log.InfoCtx(r.Context(), "retrieved codelist editions", log.Data{"code_list_id": id})
}

func (c *CodeListAPI) getEdition(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	editionModel, err := c.store.GetEdition(r.Context(), id, edition)
	if err != nil {
		handleError(ctx, w, err, data)
		return
	}

	b, err := json.Marshal(editionModel)
	if err != nil {
		handleError(ctx, w, err, data)
		return
	}

	writeBody(ctx, w, b)
	log.InfoCtx(r.Context(), "retrieved codelist edition", log.Data{"code_list_id": id, edition: edition})
}
