package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getCodeLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filterBy := r.URL.Query().Get("type")

	log.InfoCtx(ctx, "attempting to get code lists", log.Data{"type": filterBy})

	codeLists, err := c.store.GetCodeLists(r.Context(), filterBy)
	if err != nil {
		log.InfoCtx(ctx, "failed to get code lists from graph", log.Data{"err": err, "type": filterBy})
		handleError(ctx, w, err, nil)
		return
	}

	for _, item := range codeLists.Items {
		if err := item.UpdateLinks(c.apiURL); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeLists endpoint: links could not be created"), nil)
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
	}

	count := len(codeLists.Items)
	codeLists.Count = count
	codeLists.Limit = count
	codeLists.TotalCount = count

	b, err := json.Marshal(codeLists)
	if err != nil {
		log.InfoCtx(ctx, "failed to marshal code lists", log.Data{"err": err})
		handleError(ctx, w, err, nil)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeLists endpoint: failed to write bytes to response"), nil)
		return
	}
	log.InfoCtx(ctx, "retrieved all codelists", nil)
}

func (c *CodeListAPI) getCodeList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	data := log.Data{"code_list_id": id}

	codeList, err := c.store.GetCodeList(ctx, id)
	if err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeList endpoint: store.GetCodeList returned an error"), data)
		handleError(ctx, w, err, nil)
		return
	}

	if err := codeList.UpdateLinks(c.apiURL); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeList endpoint: links could not be created"), nil)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(codeList)
	if err != nil {
		log.InfoCtx(ctx, "failed to marshal code list", log.Data{"err": err})
		handleError(ctx, w, err, nil)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeList endpoint: failed to write bytes to response"), data)
		return
	}

	log.InfoCtx(ctx, "getCodeList endpoint: request successful", data)
}
