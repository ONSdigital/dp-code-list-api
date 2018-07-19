package api

import (
	"encoding/json"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"net/http"
)

var codesNotFoundErrors = map[error]bool{
	datastore.ErrEditionNotFound: true,
	datastore.ErrCodeNotFound:    true,
	datastore.ErrCodesNotFound:   true,
}

func (c *CodeListAPI) getCodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}

	log.InfoCtx(ctx, "getCodes endpoint: attempting to get edition codes", data)

	codes, err := c.store.GetCodes(ctx, id, edition)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodes endpoint: store.GetCode returned an error"), data)
		if codesNotFoundErrors[err] {
			http.Error(w, err.Error(), http.StatusNotFound)
		} else {
			http.Error(w, internalServerErr, http.StatusInternalServerError)
		}
		return
	}
	b, err := json.Marshal(codes)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodes endpoint: failed to marshal codes to json bytes"), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodes endpoint: failed to write bytes to response"), nil)
		return
	}
	log.InfoCtx(ctx, "getCodes endpoint: request successful", data)
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
		if codesNotFoundErrors[err] {
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

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCode endpoint: failed to write bytes to response"), nil)
		return
	}
	log.InfoCtx(ctx, "getCode endpoint: request successful", data)
}
