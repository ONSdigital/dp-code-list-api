package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		Limit:      1,
		TotalCount: 1,
	}
)

func TestGetEditions(t *testing.T) {

	Convey("Get code list editions returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080/code-lists/%s/editions", codeListID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbEditions, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		payload, err := ioutil.ReadAll(w.Body)
		So(err, ShouldBeNil)
		apiEditions := models.Editions{}
		json.Unmarshal(payload, &apiEditions)
		So(apiEditions, ShouldResemble, expectedEditions)
	})

	Convey("Get code list editions returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})

	Convey("Get code list editions returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionsFunc: func(ctx context.Context, f string) (*dbmodels.Editions, error) {
				return &dbmodels.Editions{}, InternalError
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}

func TestGetEdition(t *testing.T) {
	Convey("Get code list edition returns a status of http ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("http://localhost:8080/code-lists/%s/editions/%s", codeListID, editionID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbEdition, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		payload, err := ioutil.ReadAll(w.Body)
		So(err, ShouldBeNil)
		apiEdition := models.Edition{}
		json.Unmarshal(payload, &apiEdition)
		So(apiEdition, ShouldResemble, expectedEdition)
	})
	Convey("Get code list edition returns a status of http not found if code list doesn't exist", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, driver.ErrNotFound
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusNotFound)
	})
	Convey("Get code list edition returns a status internal server error if store returns any other error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/12345/editions/2016", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetEditionFunc: func(ctx context.Context, f, e string) (*dbmodels.Edition, error) {
				return &dbmodels.Edition{}, InternalError
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
