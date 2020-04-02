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

// CodeList models for testing
var (
	dbCodeList        = dbmodels.CodeList{ID: codeListID}
	dbCodeListResults = dbmodels.CodeListResults{Items: []dbmodels.CodeList{dbCodeList}}

	expectedCodeList = models.CodeList{
		Links: &models.CodeListLink{
			Self:     &models.Link{ID: codeListID, Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID)},
			Editions: &models.Link{ID: "", Href: fmt.Sprintf("%s/code-lists/%s/editions", codeListURL, codeListID)},
		},
	}
	expectedCodeListResults = models.CodeListResults{
		Items:      []models.CodeList{expectedCodeList},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 1,
	}
)

func TestGetCodeLists(t *testing.T) {
	t.Parallel()

	Convey("Get code lists returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbCodeListResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeListResults{}, &expectedCodeListResults)
	})

	Convey("Get code lists returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, f string) (*dbmodels.CodeListResults, error) {
				return nil, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetCodeList(t *testing.T) {
	t.Parallel()

	Convey("Get a code list returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				So(id, ShouldEqual, codeListID)
				return &dbCodeList, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeList{}, &expectedCodeList)
	})

	Convey("Get a code list returns a status of not found", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/123", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Get a code list returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/123", codeListURL), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
