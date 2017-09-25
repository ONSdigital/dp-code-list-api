package api

import (
	"github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
	"errors"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

var INTERNAL_ERROR = errors.New("internal error")

func TestGetCodeLists(t *testing.T) {
	t.Parallel()
	Convey("Get code lists returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeListsFunc: func() (*models.CodeListResults, error) {
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
			GetCodeListsFunc: func() (*models.CodeListResults, error) {
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
			GetCodeListFunc: func(id string) (*models.CodeList, error) {
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
			GetCodeListFunc: func(id string) (*models.CodeList, error) {
				return nil, datastore.NOT_FOUND
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
			GetCodeListFunc: func(id string) (*models.CodeList, error) {
				return nil, INTERNAL_ERROR
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
