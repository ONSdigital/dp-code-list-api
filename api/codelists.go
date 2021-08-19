package api

import (
	"encoding/json"
	"net/http"
	"sort"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
	"github.com/ONSdigital/log.go/v2/log"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func (c *CodeListAPI) getCodeLists(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	filterBy := r.URL.Query().Get("type")
	logData := log.Data{}
	offsetParameter := r.URL.Query().Get("offset")
	limitParameter := r.URL.Query().Get("limit")
	offset := c.defaultOffset
	limit := c.defaultLimit
	var err error

	log.Info(ctx, "attempting to get code lists", log.Data{"type": filterBy})

	if offsetParameter != "" {
		logData["offset"] = offsetParameter
		offset, err = ValidatePositiveInt(offsetParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: offset", err, logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limitParameter != "" {
		logData["limit"] = limitParameter
		limit, err = ValidatePositiveInt(limitParameter)
		if err != nil {
			log.Error(ctx, "invalid query parameter: limit", err, logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limit > c.maxLimit {
		logData["max_limit"] = c.maxLimit
		err = errors.New("limit is greater than the maximum allowed")
		log.Error(ctx, "invalid query parameter: limit, maximum limit reached", err, logData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
			log.Error(ctx, "error updating links", errors.WithMessage(err, "getCodeLists endpoint: links could not be created"))
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
		return
	}
	log.Info(ctx, "retrieved all codelists")
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
		log.Error(ctx, "error updating links", errors.WithMessage(err, "getCodeList endpoint: links could not be created"))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(codeList)
	if err != nil {
		handleError(ctx, "failed to marshal code list", log.Data{}, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Error(ctx, "error writting body", errors.WithMessage(err, "getCodeList endpoint: failed to write bytes to response"), data)
		return
	}

	log.Info(ctx, "getCodeList endpoint: request successful", data)
}

func codelistsSlice(full []dbmodels.CodeList, offset, limit int) (sliced []dbmodels.CodeList) {
	end := offset + limit
	if end > len(full) {
		end = len(full)
	}

	if offset > len(full) || limit == 0 {
		return []dbmodels.CodeList{}
	}
	return full[offset:end]
}
