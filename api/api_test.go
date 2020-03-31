package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

var INTERNAL_ERROR = errors.New("internal error")

func TestGetCodeLists(t *testing.T) {
	t.Parallel()
	Convey("Get code lists returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, filterBy string) (*dbmodels.CodeListResults, error) {
				return &dbmodels.CodeListResults{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code lists returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, f string) (*dbmodels.CodeListResults, error) {
				return nil, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetCodeList(t *testing.T) {
	t.Parallel()
	Convey("Get a code list returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/123", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return &dbmodels.CodeList{ID: id}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get a code list returns a status of not found", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/123", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get a code list returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/123", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*dbmodels.CodeList, error) {
				return nil, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEditions(t *testing.T) {
	Convey("Get code list editions returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code list editions returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get code list editions returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEdition(t *testing.T) {
	Convey("Get code list edition returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{
					ID:    "edition",
					Label: "label",
				}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code list edition returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get code list edition returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, "", "")
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
