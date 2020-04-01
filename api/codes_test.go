package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/dp-graph/v2/graph/driver"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	getCodesErr    = errors.New("get codes error")
	testCodelistID = "$codelist_id$"
	testEdition    = "$edition$"
	testCode       = "666"

	failWriteBody = func(w http.ResponseWriter, bytes []byte) error {
		http.Error(w, internalServerErr, http.StatusInternalServerError)
		return getCodesErr
	}

	// Code DB model
	codeDbModel = dbmodels.Code{
		ID:    "ignored",
		Code:  "666",
		Label: "test",
	}

	// Code API model with updated links
	codeModel = models.Code{
		ID:    "666", // number of the beast \m/
		Label: "test",
		Links: &models.CodeLinks{
			Self: &models.Link{
				ID:   "666",
				Href: "/code-lists/$codelist_id$/editions/$edition$/codes/666",
			},
			Datasets: &models.Link{
				ID:   "",
				Href: "/code-lists/$codelist_id$/editions/$edition$/codes/666/datasets",
			},
			CodeList: &models.Link{
				ID:   "",
				Href: "/code-lists/$codelist_id$",
			},
		},
	}
)

func TestGetCodes_DatastoreError(t *testing.T) {
	Convey("Given datastore.GetCodes returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return nil, getCodesErr
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, "", "")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes", nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, testEdition)
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
			CreateCodeListAPI(router, mockDatastore, "", "")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes", nil)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, testEdition)
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

			api := CreateCodeListAPI(router, mockDatastore, "", "")
			api.writeBody = failWriteBody

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes", nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, testEdition)
			})
		})
	})
}

func TestGetCodes_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		dbResult := dbmodels.CodeResults{
			Items: []dbmodels.Code{codeDbModel},
		}
		expectedResult := &models.CodeResults{
			Count:      1,
			Limit:      1,
			TotalCount: 1,
			Items:      []models.Code{codeModel},
		}

		mockDatastore := &storetest.DataStoreMock{
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID string) (*dbmodels.CodeResults, error) {
				return &dbResult, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, "", "")

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes", nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				var res models.CodeResults
				err := json.Unmarshal(w.Body.Bytes(), &res)
				So(err, ShouldBeNil)

				So(&res, ShouldResemble, expectedResult)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, testEdition)
			})
		})
	})
}

func TestGetCode_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return &codeDbModel, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, "", "")

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes/666", nil)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				var res models.Code
				err := json.Unmarshal(w.Body.Bytes(), &res)
				So(err, ShouldBeNil)

				So(res, ShouldResemble, codeModel)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, testEdition)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, testCode)
			})
		})
	})
}

func TestGetCode_DatastoreError(t *testing.T) {
	Convey("Given datastore.GetCode returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return nil, getCodesErr
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, "", "")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes/666", nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, testEdition)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, testCode)
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
			CreateCodeListAPI(router, mockDatastore, "", "")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes/666", nil)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, testEdition)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, testCode)
			})
		})
	})
}

func TestGetCode_WriteBodyError(t *testing.T) {
	Convey("Given datastore.GetCode returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID string, codeID string) (*dbmodels.Code, error) {
				return &codeDbModel, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			api := CreateCodeListAPI(router, mockDatastore, "", "")
			api.writeBody = failWriteBody

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, "http://localhost:8080/code-lists/$codelist_id$/editions/$edition$/codes/666", nil)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, testCodelistID)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, testEdition)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, testCode)
			})
		})
	})
}
