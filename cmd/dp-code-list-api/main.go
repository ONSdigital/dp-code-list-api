package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/ONSdigital/dp-code-list-api/api"
	"github.com/ONSdigital/dp-code-list-api/config"
	"github.com/ONSdigital/dp-code-list-api/mongo"
	"github.com/ONSdigital/dp-code-list-api/stubs"
	"github.com/ONSdigital/go-ns/log"
	"github.com/ONSdigital/go-ns/server"
	"github.com/gorilla/mux"
)

func main() {
	log.Namespace = "dp-code-list-api"
	cfg, err := config.Get()
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	mongoDatastore, err := mongo.CreateMongoDataStore(cfg.MongoDBURL)
	if err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
	router := mux.NewRouter()
	_ = api.CreateCodeListAPI(router, mongoDatastore)
	log.Info("code List API listening .....", log.Data{"BIND_ADDR": cfg.BindAddr})
	srv := server.New(cfg.BindAddr, router)
	if err := srv.ListenAndServe(); err != nil {
		log.Error(err, nil)
		os.Exit(1)
	}
}

//func main() {
//	log.Namespace = "dp-code-list-api"
//	configuration, configErr := config.Get()
//	if configErr != nil {
//		log.Error(configErr, nil)
//		os.Exit(1)
//	}
//
//	router := mux.NewRouter()
//	router.Path("/code-lists").HandlerFunc(codeListsHandler)
//	router.Path("/code-lists/{codelist_id}").HandlerFunc(codeListHandler)
//	router.Path("/code-lists/{codelist_id}/codes").HandlerFunc(codesHandler)
//	router.Path("/code-lists/{codelist_id}/codes/{code_id}").HandlerFunc(codeHandler)
//
//	log.Debug("starting http server", log.Data{"bind_addr": configuration.BindAddr})
//	srv := server.New(configuration.BindAddr, router)
//	if err := srv.ListenAndServe(); err != nil {
//		log.Error(err, nil)
//		os.Exit(1)
//	}
//}

type codeListsResponse struct {
	Count        int                `json:"count"`
	Start        int                `json:"start_index"`
	ItemsPerPage int                `json:"items_per_page"`
	Items        []codeListResponse `json:"items"`
	TotalCount   int                `json:"total_count"`
}

type codeListResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	CodeListURL string `json:"code_list_url"`
	CodesURL    string `json:"codes_url"`
}

type codeListCodesResponse struct {
	Count        int            `json:"count"`
	Start        int            `json:"start_index"`
	ItemsPerPage int            `json:"items_per_page"`
	Items        []codeListCode `json:"items"`
	TotalCount   int            `json:"total_count"`
}

type codeListCode struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Label   string `json:"label"`
	CodeURL string `json:"code_url"`
}

var codeListsAPIResponse = codeListsResponse{
	Items:        []codeListResponse{timeCodeListAPIResponse, goodsAndServicesCodeListAPIResponse},
	Count:        2,
	TotalCount:   2,
	ItemsPerPage: 10,
	Start:        0,
}

// FIXME use fully qualified URL for CodesURL
var timeCodeListAPIResponse = codeListResponse{
	ID:          stubs.Time.ID,
	Name:        stubs.Time.Name,
	CodeListURL: "/code-lists/" + stubs.Time.ID,
	CodesURL:    "/code-lists/" + stubs.Time.ID + "/codes",
}

// FIXME use fully qualified URL for CodesURL
var goodsAndServicesCodeListAPIResponse = codeListResponse{
	ID:          stubs.GoodsAndServicesIndices.ID,
	Name:        stubs.GoodsAndServicesIndices.Name,
	CodeListURL: "/code-lists/" + stubs.GoodsAndServicesIndices.ID,
	CodesURL:    "/code-lists/" + stubs.GoodsAndServicesIndices.ID + "/codes",
}

var timeCodeListCodesAPIResponse = func() codeListCodesResponse {
	r := codeListCodesResponse{
		Count:        len(stubs.Time.Codes),
		TotalCount:   len(stubs.Time.Codes),
		ItemsPerPage: len(stubs.Time.Codes),
		Start:        0,
	}

	for _, i := range stubs.Time.Codes {
		r.Items = append(r.Items, codeListCode{
			ID:      i.ID,
			Code:    i.Code,
			Label:   i.Label,
			CodeURL: "/code-lists/" + stubs.Time.ID + "/codes/" + i.ID,
		})
	}

	return r
}()

var goodsAndServicesCodeListCodesAPIResponse = func() codeListCodesResponse {
	r := codeListCodesResponse{
		Count:        len(stubs.GoodsAndServicesIndices.Codes),
		TotalCount:   len(stubs.GoodsAndServicesIndices.Codes),
		ItemsPerPage: len(stubs.GoodsAndServicesIndices.Codes),
		Start:        0,
	}

	for _, i := range stubs.GoodsAndServicesIndices.Codes {
		r.Items = append(r.Items, codeListCode{
			ID:      i.ID,
			Code:    i.Code,
			Label:   i.Label,
			CodeURL: "/code-lists/" + stubs.GoodsAndServicesIndices.ID + "/codes/" + i.ID,
		})
	}

	return r
}()

var codeListStubMap = map[string]stubs.CodeList{
	stubs.GoodsAndServicesIndices.ID: stubs.GoodsAndServicesIndices,
	stubs.Time.ID:                    stubs.Time,
}

var codeListsMap = map[string]codeListResponse{
	stubs.GoodsAndServicesIndices.ID: goodsAndServicesCodeListAPIResponse,
	stubs.Time.ID:                    timeCodeListAPIResponse,
}

var codeListCodesMap = map[string]codeListCodesResponse{
	stubs.GoodsAndServicesIndices.ID: goodsAndServicesCodeListCodesAPIResponse,
	stubs.Time.ID:                    timeCodeListCodesAPIResponse,
}

func codeListsHandler(w http.ResponseWriter, req *http.Request) {
	b, err := json.Marshal(&codeListsAPIResponse)
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func codeListHandler(w http.ResponseWriter, req *http.Request) {
	codeListID := mux.Vars(req)["codelist_id"]

	if _, ok := codeListsMap[codeListID]; !ok {
		log.DebugR(req, "code list not found", log.Data{"code_list_id": codeListID})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(codeListsMap[codeListID])
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func codesHandler(w http.ResponseWriter, req *http.Request) {
	codeListID := mux.Vars(req)["codelist_id"]

	if _, ok := codeListsMap[codeListID]; !ok {
		log.DebugR(req, "code list not found", log.Data{"code_list_id": codeListID})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, err := json.Marshal(codeListCodesMap[codeListID])
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func codeHandler(w http.ResponseWriter, req *http.Request) {
	codeListID := mux.Vars(req)["codelist_id"]
	codeID := mux.Vars(req)["code_id"]

	if _, ok := codeListsMap[codeListID]; !ok {
		log.DebugR(req, "code list not found", log.Data{"code_list_id": codeListID})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, ok := codeListStubMap[codeListID].Codes[codeID]; !ok {
		log.DebugR(req, "code not found in list", log.Data{"code_list_id": codeListID, "code_id": codeID})
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cl := codeListStubMap[codeListID].Codes[codeID]
	b, err := json.Marshal(codeListCode{
		ID:      cl.ID,
		Code:    cl.Code,
		Label:   cl.Label,
		CodeURL: "/code-lists/" + codeListID + "/codes/" + codeID,
	})
	if err != nil {
		log.ErrorR(req, err, nil)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
