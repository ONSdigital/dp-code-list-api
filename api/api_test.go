package api

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	"github.com/ONSdigital/dp-code-list-api/models"
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
			GetCodeListsFunc: func(ctx context.Context, f string) (*models.CodeListResults, error) {
				return &models.CodeListResults{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code lists returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func(ctx context.Context, f string) (*models.CodeListResults, error) {
				return nil, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
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
			GetCodeListFunc: func(ctx context.Context, id string) (*models.CodeList, error) {
				return &models.CodeList{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get a code list returns a status of not found", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/123", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*models.CodeList, error) {
				return nil, datastore.ErrCodeListNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get a code list returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/123", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListFunc: func(ctx context.Context, id string) (*models.CodeList, error) {
				return nil, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEditions(t *testing.T) {
	Convey("Get code list editions returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*models.Editions, error) {
				return &models.Editions{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code list editions returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*models.Editions, error) {
				return &models.Editions{}, datastore.ErrEditionsNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get code list editions returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*models.Editions, error) {
				return &models.Editions{}, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEdition(t *testing.T) {
	Convey("Get code list edition returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*models.Edition, error) {
				return &models.Edition{}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)
	})
	Convey("Get code list edition returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*models.Edition, error) {
				return &models.Edition{}, datastore.ErrEditionNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get code list edition returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*models.Edition, error) {
				return &models.Edition{}, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
