package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/pkg/errors"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// CodeList models for testing
var (
	dbCodeList1       = dbmodels.CodeList{ID: codeListID1}
	dbCodeList2       = dbmodels.CodeList{ID: codeListID2}
	dbCodeListResults = dbmodels.CodeListResults{Items: []dbmodels.CodeList{dbCodeList1, dbCodeList2}}

	expectedCodeList1 = models.CodeList{
		Links: &models.CodeListLink{
			Self:     &models.Link{ID: codeListID1, Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1)},
			Editions: &models.Link{ID: "", Href: fmt.Sprintf("%s/code-lists/%s/editions", codeListURL, codeListID1)},
		},
	}

	expectedCodeList2 = models.CodeList{
		Links: &models.CodeListLink{
			Self:     &models.Link{ID: codeListID2, Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID2)},
			Editions: &models.Link{ID: "", Href: fmt.Sprintf("%s/code-lists/%s/editions", codeListURL, codeListID2)},
		},
	}

	expectedCodeListResults = models.CodeListResults{
		Items:      []models.CodeList{expectedCodeList1, expectedCodeList2},
		Count:      2,
		Offset:     0,
		Limit:      20,
		TotalCount: 2,
	}

	paginationTestOne = models.CodeListResults{
		Items:      []models.CodeList{expectedCodeList1},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 2,
	}

	paginationTestTwo = models.CodeListResults{
		Items:      []models.CodeList{},
		Count:      0,
		Offset:     1,
		Limit:      7,
		TotalCount: 2,
	}

	paginationTestFour = models.CodeListResults{
		Items:      nil,
		Count:      0,
		Offset:     3,
		Limit:      1,
		TotalCount: 2,
	}
)

func TestGetCodeLists(t *testing.T) {
	t.Parallel()

	Convey("Get code lists without query parameters returns a status of Ok with the default offset and limit values", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbCodeListResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeListResults{}, &expectedCodeListResults)
	})

	Convey("When valid limit and offset query parameters are provided, then return codelist information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?offset=0&limit=1", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbCodeListResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeListResults{}, &paginationTestOne)
	})

	Convey("When offset value greater than count provided, then return zero items", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?offset=3&limit=1", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbCodeListResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeListResults{}, &paginationTestFour)
	})

	Convey("When datastore GetCodeLists returns an error, then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		testError := errors.New("datastore.GtCodeLists failed")
		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return nil, testError
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When datastore GetCodeLists returns an invalid codeList (no ID), then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbmodels.CodeListResults{
					Items: []dbmodels.CodeList{{ID: ""}},
				}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When write response body returns an error, then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbCodeListResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.writeBody = failWriteBody
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When a negative offset query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?offset=-1", codeListURL), nil)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When a negative limit query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?limit=-1", codeListURL), nil)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When limit above default maximum is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?limit=1001", codeListURL), nil)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)

	})

	Convey("When non-integer query parameter value provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists?offset=x&limit=y", codeListURL), nil)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)

	})

}

func TestGetCodeList(t *testing.T) {
	t.Parallel()

	Convey("Get a code list returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				So(id, ShouldEqual, codeListID1)
				return &dbCodeList1, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeList{}, &expectedCodeList1)
	})

	Convey("Get a code list returns a status of not found", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("When datastore GetCodeList returns an invalid codeList (no ID), then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, code string) (*dbmodels.CodeList, error) {
				return &dbmodels.CodeList{ID: ""}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When datastore GetCodeList returns an error, then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When write response body returns an error, then 500 status is returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				So(id, ShouldEqual, codeListID1)
				return &dbCodeList1, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.writeBody = failWriteBody
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
