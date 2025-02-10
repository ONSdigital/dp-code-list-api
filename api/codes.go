package api

import (
	"encoding/json"
	"net/http"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
	"github.com/ONSdigital/dp-net/v2/links"
	"github.com/ONSdigital/log.go/v2/log"
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

	log.Info(ctx, "getCodes endpoint: attempting to get edition codes", data)

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

	totalCount, err := c.store.CountCodes(ctx, id, edition)
	if err != nil {
		handleError(ctx, "getCodes endpoint: store.CountCodes returned an error", data, err, w)
		return
	}

	dbCodes := &dbmodels.CodeResults{}

	if limit > 0 && totalCount > 0 {
		dbCodes, err = c.store.GetCodes(ctx, id, edition)
		if err != nil {
			handleError(ctx, "getCodes endpoint: store.GetCodes returned an error", data, err, w)
			return
		}
	}

	slicedResults := codesSlice(dbCodes.Items, offset, limit)

	codes := models.NewCodeResults(slicedResults)

	codeListLinksBuilder := links.FromHeadersOrDefault(&r.Header, c.parsedAPIURL)

	for i, item := range codes.Items {
		if err := item.UpdateLinks(c.apiURL, id, edition); err != nil {
			log.Error(ctx, "error updating links", errors.WithMessage(err, "getCodes endpoint: links could not be created"))
			http.Error(w, internalServerErr, http.StatusInternalServerError)
			return
		}

		if c.enableURLRewriting {
			item.Links.Self.Href, err = codeListLinksBuilder.BuildLink(item.Links.Self.Href)
			if err != nil {
				log.Error(ctx, "could not build self link", err, log.Data{"link": item.Links.Self.Href})
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			item.Links.CodeList.Href, err = codeListLinksBuilder.BuildLink(item.Links.CodeList.Href)
			if err != nil {
				log.Error(ctx, "could not build code list link", err, log.Data{"link": item.Links.CodeList.Href})
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			item.Links.Datasets.Href, err = codeListLinksBuilder.BuildLink(item.Links.Datasets.Href)
			if err != nil {
				log.Error(ctx, "could not build datasets link", err, log.Data{"link": item.Links.Datasets.Href})
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		codes.Items[i] = item
	}

	count := len(slicedResults)
	codes.Count = count
	codes.Offset = offset
	codes.Limit = limit
	codes.TotalCount = int(totalCount)

	b, err := json.Marshal(codes)
	if err != nil {
		handleError(ctx, "failed to marshal codes", data, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		return
	}

	log.Info(ctx, "getCodes endpoint: request successful", data)
}

func (c *CodeListAPI) getCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	id := vars["id"]
	edition := vars["edition"]
	code := vars["code"]
	data := log.Data{"codelist_id": id, "edition": edition, "code": code}

	log.Info(ctx, "getCode: attempting to get code list code", data)

	dbCode, err := c.store.GetCode(ctx, id, edition, code)
	if err != nil {
		handleError(ctx, "getCode endpoint: store.GetCode returned an error", data, err, w)
		return
	}
	apiCode := models.NewCode(dbCode)

	if err := apiCode.UpdateLinks(c.apiURL, id, edition); err != nil {
		log.Error(ctx, "error updating links", errors.WithMessage(err, "getCode endpoint: links could not be created"))
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return
	}

	codeListLinksBuilder := links.FromHeadersOrDefault(&r.Header, c.parsedAPIURL)

	if c.enableURLRewriting {
		apiCode.Links.Self.Href, err = codeListLinksBuilder.BuildLink(apiCode.Links.Self.Href)
		if err != nil {
			log.Error(ctx, "could not build self link", err, log.Data{"link": apiCode.Links.Self.Href})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiCode.Links.CodeList.Href, err = codeListLinksBuilder.BuildLink(apiCode.Links.CodeList.Href)
		if err != nil {
			log.Error(ctx, "could not build code list link", err, log.Data{"link": apiCode.Links.CodeList.Href})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		apiCode.Links.Datasets.Href, err = codeListLinksBuilder.BuildLink(apiCode.Links.Datasets.Href)
		if err != nil {
			log.Error(ctx, "could not build datasets link", err, log.Data{"link": apiCode.Links.Datasets.Href})
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	b, err := json.Marshal(apiCode)
	if err != nil {
		handleError(ctx, "failed to marshal code", data, err, w)
		return
	}

	if err := c.writeBody(w, b); err != nil {
		log.Error(ctx, "error writting body", errors.WithMessage(err, "getCode endpoint: failed to write bytes to response"))
		return
	}

	log.Info(ctx, "getCode endpoint: request successful", data)
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
