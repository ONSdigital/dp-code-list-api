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

func (c *CodeListAPI) getCodes(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	data := log.Data{"codelist_id": id, "edition": edition}
	logData := log.Data{}
	offsetParameter := r.URL.Query().Get("offset")
	limitParameter := r.URL.Query().Get("limit")
	offset := c.defaultOffset
	limit := c.defaultLimit
	var err error

	log.Event(ctx, "getCodes endpoint: attempting to get edition codes", log.INFO, data)

	if offsetParameter != "" {
		logData["offset"] = offsetParameter
		offset, err = ValidatePositiveInt(offsetParameter)
		if err != nil {
			log.Event(ctx, "invalid query parameter: offset", log.ERROR, log.Error(err), logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limitParameter != "" {
		logData["limit"] = limitParameter
		limit, err = ValidatePositiveInt(limitParameter)
		if err != nil {
			log.Event(ctx, "invalid query parameter: limit", log.ERROR, log.Error(err), logData)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if limit > c.maxLimit {
		logData["max_limit"] = c.maxLimit
		err = errors.New("limit is greater than the maximum allowed")
		log.Event(ctx, "invalid query parameter: limit, maximum limit reached", log.ERROR, log.Error(err), logData)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbCodes, err := c.store.GetCodes(ctx, id, edition)
	if err != nil {
		handleError(ctx, "getCodes endpoint: store.GetCode returned an error", data, err, w)
		return
	}

	totalCount := len(dbCodes.Items)

	sort.Slice(dbCodes.Items, func(i, j int) bool {
		return dbCodes.Items[i].ID < dbCodes.Items[j].ID
	})

	slicedResults := codesSlice(dbCodes.Items, offset, limit)

	codes := models.NewCodeResults(slicedResults)

	for i, item := range codes.Items {
		if err := item.UpdateLinks(c.apiURL, id, edition); err != nil {
			log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: links could not be created")))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}
		codes.Items[i] = item
	}

	count := len(slicedResults)
	codes.Count = count
	codes.Offset = offset
	codes.Limit = limit
	codes.TotalCount = totalCount

	b, err := json.Marshal(codes)
	if err != nil {
		log.Event(ctx, "marshal error", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: failed to marshal codes to json bytes")), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCodes endpoint: failed to write bytes to response")))
		return
	}

	log.Event(ctx, "getCodes endpoint: request successful", log.INFO, data)
}

func (c *CodeListAPI) getCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	data := log.Data{"codelist_id": id, "edition": edition, "code": code}

	log.Event(ctx, "getCode: attempting to get code list code", log.INFO, data)

	dbCode, err := c.store.GetCode(ctx, id, edition, code)
	if err != nil {
		handleError(ctx, "getCode endpoint: store.GetCode returned an error", data, err, w)
		return
	}
	apiCode := models.NewCode(dbCode)

	if err := apiCode.UpdateLinks(c.apiURL, id, edition); err != nil {
		log.Event(ctx, "error updating links", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: links could not be created")))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(apiCode)
	if err != nil {
		log.Event(ctx, "marshal error", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: error attempting to marshal result to JSON")), data)
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Event(ctx, "error writting body", log.ERROR, log.Error(errors.WithMessage(err, "getCode endpoint: failed to write bytes to response")))
		return
	}

	log.Event(ctx, "getCode endpoint: request successful", log.INFO, data)
}

func codesSlice(full []dbmodels.Code, offset, limit int) (sliced []dbmodels.Code) {
	end := offset + limit
	if end > len(full) {
		end = len(full)
	}

	if offset > len(full) || limit == 0 {
		return []dbmodels.Code{}
	}
	return full[offset:end]
}
