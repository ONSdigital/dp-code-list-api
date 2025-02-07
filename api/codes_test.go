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
	dbCode1 = dbmodels.Code{
		ID:    "ignored",
		Code:  codeID1,
		Label: "test one",
	}

	dbCode2 = dbmodels.Code{
		ID:    "ignored",
		Code:  codeID2,
		Label: "test two",
	}

	dbCodeResults = dbmodels.CodeResults{
		Items: []dbmodels.Code{dbCode1, dbCode2},
	}

	expectedCode1 = models.Code{
		ID:    codeID1,
		Label: "test one",
		Links: &models.CodeLinks{
			Self: &models.Link{
				ID:   codeID1,
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID1),
			},
			Datasets: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID1, editionID1, codeID1),
			},
			CodeList: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1),
			},
		},
	}

	expectedCode2 = models.Code{
		ID:    codeID2,
		Label: "test two",
		Links: &models.CodeLinks{
			Self: &models.Link{
				ID:   codeID2,
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID2),
			},
			Datasets: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID1, editionID1, codeID2),
			},
			CodeList: &models.Link{
				ID:   "",
				Href: fmt.Sprintf("%s/code-lists/%s", codeListURL, codeListID1),
			},
		},
	}

	codePaginationTestZero = models.CodeResults{
		Items:      nil,
		Count:      0,
		Offset:     0,
		Limit:      0,
		TotalCount: 2,
	}

	codePaginationTestOne = models.CodeResults{
		Items:      []models.Code{expectedCode1},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 2,
	}

	codePaginationTestThree = models.CodeResults{
		Items:      nil,
		Count:      0,
		Offset:     3,
		Limit:      1,
		TotalCount: 2,
	}
)

func TestGetCodes_DatastoreError(t *testing.T) {
	Convey("Given datastore.CountCodes returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 0, ErrInternal
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)
			})
		})
	})

	Convey("Given datastore.GetCodes returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 1, nil
			},
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID1 string) (*dbmodels.CodeResults, error) {
				return nil, ErrInternal
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID1)
			})
		})
	})
}

func TestGetCodes_EditionNotFound(t *testing.T) {
	Convey("Given datastore.CountCodes returns an edition not found error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 1, driver.ErrNotFound
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)
			})
		})
	})
}

func TestGetCodes_InvalidCode(t *testing.T) {
	Convey("Given GetCodes returns an invalid code (empty ID)", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 1, nil
			},
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID1 string) (*dbmodels.CodeResults, error) {
				return &dbmodels.CodeResults{
					Items: []dbmodels.Code{{ID: ""}},
				}, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)

				So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID1)
			})
		})
	})
}

func TestGetCodes_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		dbResult := dbmodels.CodeResults{
			Items: []dbmodels.Code{dbCode1, dbCode2},
		}
		expectedResult := models.CodeResults{
			Count:      2,
			Limit:      20,
			TotalCount: 2,
			Items:      []models.Code{expectedCode1, expectedCode2},
		}

		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return int64(2), nil
			},
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID1 string) (*dbmodels.CodeResults, error) {
				return &dbResult, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				validateBody(w.Body, &models.CodeResults{}, &expectedResult)

				Convey("and CountCodes and GetCodes are called with the expected paramters", func() {
					So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
					So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
					So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)

					So(mockDatastore.GetCodesCalls(), ShouldHaveLength, 1)
					So(mockDatastore.GetCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
					So(mockDatastore.GetCodesCalls()[0].EditionID, ShouldEqual, editionID1)
				})
			})
		})
	})

	Convey("Given a valid request for a codelist without codes", t, func() {
		expectedResult := models.CodeResults{
			Count:      0,
			Limit:      20,
			TotalCount: 0,
			Items:      nil,
		}

		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return int64(0), nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes", codeListURL, codeListID1, editionID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				validateBody(w.Body, &models.CodeResults{}, &expectedResult)

				Convey("and CountCodes is called with the expected paramters", func() {
					So(mockDatastore.CountCodesCalls(), ShouldHaveLength, 1)
					So(mockDatastore.CountCodesCalls()[0].CodeListID, ShouldEqual, codeListID1)
					So(mockDatastore.CountCodesCalls()[0].Edition, ShouldEqual, editionID1)
				})
			})
		})
	})
}

func TestGetCodes_Pagination(t *testing.T) {
	t.Parallel()

	Convey("When valid limit and offset query parameters are provided, then return codes information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?offset=0&limit=1", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 2, nil
			},
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID1 string) (*dbmodels.CodeResults, error) {
				return &dbCodeResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeResults{}, &codePaginationTestOne)
	})

	Convey("When offset value greater than count provided, then return zero items", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?offset=3&limit=1", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 2, nil
			},
			GetCodesFunc: func(ctx context.Context, codeListID string, editionID1 string) (*dbmodels.CodeResults, error) {
				return &dbCodeResults, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeResults{}, &codePaginationTestThree)
	})

	Convey("When limit is zero, then return zero items and only call dp-graph CountCodes", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?limit=0", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			CountCodesFunc: func(ctx context.Context, codeListID string, edition string) (int64, error) {
				return 2, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.CodeResults{}, &codePaginationTestZero)
	})

	Convey("When a negative offset query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?offset=-1", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When a negative limit query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?limit=-1", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When limit above default maximum is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?limit=1001", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When non-integer query parameter value provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes?offset=x&limit=y", codeListURL, codeListID1, editionID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})
}

func TestGetCode_Success(t *testing.T) {
	Convey("Given a valid request", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID1 string, codeID string) (*dbmodels.Code, error) {
				return &dbCode1, nil
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()

			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 200 status is returned with the expected body", func() {
				So(w.Code, ShouldEqual, http.StatusOK)

				validateBody(w.Body, &models.Code{}, &expectedCode1)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID1)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID1)
			})
		})
	})
}

func TestGetCode_DatastoreError(t *testing.T) {
	Convey("Given datastore.GetCode returns an error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID1 string, codeID string) (*dbmodels.Code, error) {
				return nil, ErrInternal
			},
		}

		Convey("when getCodes is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID1)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID1)
			})
		})
	})
}

func TestGetCode_EditionNotFound(t *testing.T) {
	Convey("Given datastore.GetCode returns edition not found error", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID1 string, codeID string) (*dbmodels.Code, error) {
				return nil, driver.ErrNotFound
			},
		}

		Convey("when getCode is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 404 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusNotFound)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, driver.ErrNotFound.Error())

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID1)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID1)
			})
		})
	})
}

func TestGetCode_InvalidCode(t *testing.T) {
	Convey("Given datastore.GetCode returns an invalid code (empty ID)", t, func() {
		mockDatastore := &storetest.DataStoreMock{
			GetCodeFunc: func(ctx context.Context, codeListID string, editionID1 string, codeID string) (*dbmodels.Code, error) {
				return &dbmodels.Code{ID: ""}, nil
			},
		}

		Convey("when getCode is called", func() {
			router := mux.NewRouter()
			CreateCodeListAPI(router, mockDatastore, codeListURL, datasetURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)

			router.ServeHTTP(w, r)

			Convey("then a 500 status is returned", func() {
				So(w.Code, ShouldEqual, http.StatusInternalServerError)
				So(strings.TrimSpace(w.Body.String()), ShouldEqual, internalServerErr)

				So(mockDatastore.GetCodeCalls(), ShouldHaveLength, 1)
				So(mockDatastore.GetCodeCalls()[0].CodeListID, ShouldEqual, codeListID1)
				So(mockDatastore.GetCodeCalls()[0].EditionID, ShouldEqual, editionID1)
				So(mockDatastore.GetCodeCalls()[0].CodeID, ShouldEqual, codeID1)
			})
		})
	})
}
