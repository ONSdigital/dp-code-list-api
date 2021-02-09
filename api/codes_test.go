package api

import (
	"context"
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

	dbCodeResults = dbmodels.CodeResults{Items: []dbmodels.Code{dbCode}}

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

	codePaginationTestOne = models.CodeResults{
		Items:      []models.Code{expectedCode},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 1,
	}

	codePaginationTestTwo = models.CodeResults{
		Items:      []models.Code{expectedCode},
		Count:      0,
		Offset:     1,
		Limit:      7,
		TotalCount: 1,
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
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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

			api := CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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
		expectedResult := models.CodeResults{
			Count:      1,
			Limit:      20,
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

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID, editionID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				validateBody(w.Body, &models.CodeResults{}, &expectedResult)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID)
			})
		})
	})
}

func TestGetCodes_Pagination(t *testing.T) {
	t.Parallel()

	Convey("When valid limit and offset query parameters are provided, then return codes information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?offset=0&limit=1", codeListURL, codeListID, editionID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return &dbCodeResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeResults{}, &codePaginationTestOne)
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

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID, editionID, codeID), nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				validateBody(w.Body, &models.Code{}, &expectedCode)

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
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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
			api := CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit)
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
