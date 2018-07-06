package api

import (
	"net/http"
	"github.com/ONSdigital/go-ns/log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

var notFoundErrors = map[error]bool{
	datastore.ErrEditionNotFound: true,
	datastore.ErrCodeNotFound:    true,
}

func (c *CodeListAPI) getCodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.InfoCtx(ctx, "getting edition codes", data)

	codes, err := c.store.GetCodes(ctx, id, edition)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodes endpoint: store.GetCode returned an error"), data)
		if notFoundErrors[err] {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, internalServerErr, http.StatusInternalServerError)
		}
		return
	}
	b, _ := json.Marshal(codes)
	writeBody(ctx, w, b)
}

func (c *CodeListAPI) getCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	data := log.Data{"codelist_id": id, "edition": edition, "code": code}

	log.InfoCtx(ctx, "getCode: attempting to get code list code", data)

	result, err := c.store.GetCode(ctx, id, edition, code)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCode endpoint: store.GetCode returned an error"), data)
		if notFoundErrors[err] {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, internalServerErr, http.StatusInternalServerError)
		}
		return
	}
	b, err := json.Marshal(result)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCode endpoint: error attempting to marshal result to JSON"), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	log.InfoCtx(ctx, "getCode endpoint: request successful", data)
	writeBody(ctx, w, b)
}
