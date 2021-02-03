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

func (c *CodeListAPI) getCodeLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filterBy := r.URL.Query().Get("type")

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

	log.Event(ctx, "attempting to get code lists", log.INFO, log.Data{"type": filterBy})

	dbCodeLists, err := c.store.GetCodeLists(r.Context(), filterBy)
	if err != nil {
		handleError(ctx, "failed to get code lists from graph", log.Data{"type": filterBy}, err, w)
		return
	}

	totalCount := len(dbCodeLists.Items)

	sort.Slice(dbCodeLists.Items, func(i, j int) bool {
		return dbCodeLists.Items[i].ID < dbCodeLists.Items[j].ID
	})

	slicedResults := codelistsSlice(dbCodeLists.Items, offset, limit)

	codeLists := models.NewCodeListResults(slicedResults)

	for i, item := range codeLists.Items {
		if err := item.UpdateLinks(c.apiURL); err != nil {
			log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodeLists endpoint: links could not be created")))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
		codeLists.Items[i] = item
	}

	count := len(slicedResults)
	codeLists.Count = count
	codeLists.Offset = offset
	codeLists.Limit = limit
	codeLists.TotalCount = totalCount

	b, err := json.Marshal(codeLists)
	if err != nil {
		handleError(ctx, "failed to marshal code lists", log.Data{}, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCodeLists endpoint: failed to write bytes to response")))
		return
	}
	log.Event(ctx, "retrieved all codelists", log.INFO)
}

func (c *CodeListAPI) getCodeList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	data := log.Data{"code_list_id": id}

	dbCodeList, err := c.store.GetCodeList(ctx, id)
	if err != nil {
		handleError(ctx, "getCodeList endpoint: store.GetCodeList returned an error", data, err, w)
		return
	}

	codeList := models.NewCodeList(dbCodeList)

	if err := codeList.UpdateLinks(c.apiURL); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodeList endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(codeList)
	if err != nil {
		handleError(ctx, "failed to marshal code list", log.Data{}, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCodeList endpoint: failed to write bytes to response")), data)
		return
	}

	log.Event(ctx, "getCodeList endpoint: request successful", log.INFO, data)
}

func codelistsSlice(full []dbmodels.CodeList, offset, limit int) (sliced []dbmodels.CodeList) {
	end := offset + limit
	if limit == 0 || end > len(full) {
		end = len(full)
	}

	if offset > len(full) {
		return []dbmodels.CodeList{}
	}
	return full[offset:end]
}
