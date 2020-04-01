package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// Code models for testing
var (
	dbCode = dbmodels.Code{
		ID:    "ignored",
		Code:  codeID,
		Label: "test",
	}

	expectedCode = models.Code{
		ID:    codeID,
		Label: "test",
		Links: &models.CodeLinks{
			Self: &models.Link{
				ID:   codeID,
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID),
			},
			Datasets: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID, editionID, codeID),
			},
			CodeList: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID),
			},
		},
	}
)

var failWriteBody = func(w http.ResponseWriter, bytes []byte) error {
	http.Error(w, internalServerErr, http.StatusInternalServerError)
	return ErrInternal
}

func TestGetCodes_DatastoreError(t *testing.T) {

	Convey("Given datastore.GetCodes returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return nil, ErrInternal
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID)
			})
		})
	})
}

func TestGetCodes_EditionNotFound(t *testing.T) {
	Convey("Given datastore.GetCodes returns an edition not found error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return nil, driver.ErrNotFound
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID)
			})
		})
	})
}

func TestGetCodes_WriteBodyError(t *testing.T) {
	Convey("Given write response body returns an error ", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return &dbmodels.CodeResults{}, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			api := CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			api.writeBody = failWriteBody

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID)
			})
		})
	})
}

func TestGetCodes_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		dbResult := dbmodels.CodeResults{
			Items: []dbmodels.Code{dbCode},
		}
		expectedResult := &models.CodeResults{
			Count:      1,
			Limit:      1,
			TotalCount: 1,
			Items:      []models.Code{expectedCode},
		}

		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return &dbResult, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				var res models.CodeResults
				err := json.Unmarshal(w.Body.Bytes(), &res)
				So(err, ShouldBeNil)

				So(&res, ShouldResemble, expectedResult)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID)
			})
		})
	})
}

func TestGetCode_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return &dbCode, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				var res models.Code
				err := json.Unmarshal(w.Body.Bytes(), &res)
				So(err, ShouldBeNil)

				So(res, ShouldResemble, expectedCode)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID)
			})
		})
	})
}

func TestGetCode_DatastoreError(t *testing.T) {
	Convey("Given datastore.GetCode returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return nil, ErrInternal
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID)
			})
		})
	})
}

func TestGetCode_EditionNotFound(t *testing.T) {
	Convey("Given datastore.GetCode returns edition not found error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return nil, driver.ErrNotFound
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID)
			})
		})
	})
}

func TestGetCode_WriteBodyError(t *testing.T) {
	Convey("Given datastore.GetCode returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return &dbCode, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			api := CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL)
			api.writeBody = failWriteBody

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID)
			})
		})
	})
}
