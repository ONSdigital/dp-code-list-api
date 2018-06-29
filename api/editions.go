package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
)

func (c *CodeListAPI) getEditions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	editions, err := c.store.GetEditions(r.Context(), id)
	if err != nil {
		handleError(w, err)
		return
	}

	b, err := json.Marshal(editions)
	if err != nil {
		handleError(w, err)
		return
	}

	writeBody(w, b)
	log.InfoCtx(r.Context(), "retrieved codelist editions", log.Data{"code_list_id": id})
}

func (c *CodeListAPI) getEdition(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]

	editionModel, err := c.store.GetEdition(r.Context(), id, edition)
	if err != nil {
		handleError(w, err)
		return
	}

	b, err := json.Marshal(editionModel)
	if err != nil {
		handleError(w, err)
		return
	}

	writeBody(w, b)
	log.InfoCtx(r.Context(), "retrieved codelist edition", log.Data{"code_list_id": id, edition: edition})
}
