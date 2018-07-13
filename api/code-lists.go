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

	codeLists, err := c.store.GetCodeLists(r.Context(), filterBy)
	if err != nil {
		handleError(ctx, w, err, nil)
		return
	}

	count := len(codeLists.Items)

	codeLists.Count = count
	codeLists.Limit = count
	codeLists.TotalCount = count

	b, err := json.Marshal(codeLists)
	if err != nil {
		handleError(ctx, w, err, nil)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeLists endpoint: failed to write bytes to response"), nil)
		return
	}
	log.InfoCtx(r.Context(), "retrieved all codelists", nil)
}

func (c *CodeListAPI) getCodeList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	codeList, err := c.store.GetCodeList(r.Context(), id)
	if err != nil {
		handleError(ctx, w, err, nil)
		return
	}

	b, err := json.Marshal(codeList)
	if err != nil {
		handleError(ctx, w, err, nil)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.ErrorCtx(ctx, errors.WithMessage(err, "getCodeList endpoint: failed to write bytes to response"), log.Data{"code_list_id": id})
		return
	}
	log.InfoCtx(r.Context(), "retrieved codelist", log.Data{"code_list_id": id})
}
