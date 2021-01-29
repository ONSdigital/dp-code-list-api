package api

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	"github.com/ONSdigital/dp-code-list-api/models"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// Edition models for testing
var (
	dbEdition = dbmodels.Edition{
		ID:    editionID,
		Label: "label1",
	}
	dbEditions = dbmodels.Editions{
		Items: []dbmodels.Edition{
			dbEdition,
		},
	}

	expectedEdition = models.Edition{
		ID:    editionID,
		Label: "label1",
		Links: &models.EditionLinks{
			Self:     &models.Link{ID: editionID, Href: fmt.Sprintf("%s/code-lists/%s/editions/%s", codeListURL, codeListID, editionID)},
			Editions: &models.Link{ID: "", Href: fmt.Sprintf("%s/code-lists/%s/editions", codeListURL, codeListID)},
			Codes:    &models.Link{ID: "", Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID)},
		},
	}
	expectedEditions = models.Editions{
		Items:      []models.Edition{expectedEdition},
		Count:      1,
		Offset:     0,
		Limit:      100,
		TotalCount: 1,
	}

	editionsPaginationTestOne = models.Editions{
		Items:      []models.Edition{expectedEdition},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 1,
	}

	editionsPaginationTestTwo = models.Editions{
		Items:      []models.Edition{expectedEdition},
		Count:      0,
		Offset:     1,
		Limit:      7,
		TotalCount: 1,
	}

	editionsPaginationTestThree = models.Editions{
		Items:      []models.Edition{expectedEdition},
		Count:      1,
		Offset:     0,
		Limit:      0,
		TotalCount: 1,
	}
)

func TestGetEditions(t *testing.T) {

	Convey("Get code list editions returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions", codeListURL, codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbEditions, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Editions{}, &expectedEditions)
	})

	Convey("Get code list editions returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/12345/editions", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Get code list editions returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/12345/editions", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEdition(t *testing.T) {
	Convey("Get code list edition returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s", codeListURL, codeListID, editionID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbEdition, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Edition{}, &expectedEdition)
	})

	Convey("Get code list edition returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/12345/editions/2016", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Get code list edition returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/12345/editions/2016", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEditions_Pagination(t *testing.T) {
	t.Parallel()

	Convey("When valid limit and offset query parameters are provided, then return editions information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions?offset=0&limit=1", codeListURL, codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbEditions, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Editions{}, &editionsPaginationTestOne)
	})

	Convey("When valid limit above maximum and offset query parameters are provided, then return editions information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions?offset=1&limit=7", codeListURL, codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbEditions, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Editions{}, &editionsPaginationTestTwo)
	})

	Convey("When a negative limit and offset query parameters are provided, then return codes information with offset and limit equal to zero", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions?offset=-1&limit=-7", codeListURL, codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbEditions, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Editions{}, &editionsPaginationTestThree)
	})
}
