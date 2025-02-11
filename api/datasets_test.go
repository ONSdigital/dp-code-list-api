package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	storetest "github.com/ONSdigital/dp-code-list-api/datastore/datastoretest"
	"github.com/ONSdigital/dp-code-list-api/models"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// Dataset models for testing
var (
	dbDataset1 = dbmodels.Dataset{
		ID:             datasetID1,
		DimensionLabel: "test Dataset One",
		Editions: []dbmodels.DatasetEdition{
			{
				ID:            datasetEditionID,
				CodeListID:    codeID1,
				LatestVersion: latestDatasetEditionVersion,
			},
		},
	}

	dbDataset2 = dbmodels.Dataset{
		ID:             datasetID2,
		DimensionLabel: "test Dataset Two",
	}

	dbDatasets = dbmodels.Datasets{
		Items: []dbmodels.Dataset{dbDataset1, dbDataset2},
	}

	expectedDataset1 = models.Dataset{
		ID:             datasetID1,
		DimensionLabel: "test Dataset One",
		Editions: []models.DatasetEdition{
			{
				ID:            datasetEditionID,
				LatestVersion: latestDatasetEditionVersion,
				Links: &models.DatasetEditionLinks{
					Self:             &models.Link{ID: datasetEditionID, Href: fmt.Sprintf("%s/datasets/%s/editions/%s", datasetURL, datasetID1, datasetEditionID)},
					DatasetDimension: &models.Link{ID: codeListID1, Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d/dimensions/%s", datasetURL, datasetID1, datasetEditionID, latestDatasetEditionVersion, codeListID1)},
					LatestVersion:    &models.Link{ID: fmt.Sprintf("%d", latestDatasetEditionVersion), Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d", datasetURL, datasetID1, datasetEditionID, latestDatasetEditionVersion)},
				},
			},
		},
		Links: &models.DatasetLinks{
			Self: &models.Link{ID: datasetID1, Href: fmt.Sprintf("%s/datasets/%s", datasetURL, datasetID1)},
		},
	}

	expectedDataset2 = models.Dataset{
		ID:             datasetID2,
		DimensionLabel: "test Dataset Two",
		Links: &models.DatasetLinks{
			Self: &models.Link{ID: datasetID2, Href: fmt.Sprintf("%s/datasets/%s", datasetURL, datasetID2)},
		},
		Editions: []models.DatasetEdition{},
	}

	expectedDatasets = models.Datasets{
		Items:      []models.Dataset{expectedDataset1, expectedDataset2},
		Count:      2,
		Offset:     0,
		Limit:      20,
		TotalCount: 2,
	}

	datasetPaginationTestOne = models.Datasets{
		Items:      []models.Dataset{expectedDataset1},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 2,
	}

	datasetPaginationTestTwo = models.Datasets{
		Items:      []models.Dataset{expectedDataset2},
		Count:      1,
		Offset:     1,
		Limit:      7,
		TotalCount: 2,
	}

	datasetPaginationTestThree = models.Datasets{
		Items:      []models.Dataset{},
		Count:      0,
		Offset:     3,
		Limit:      1,
		TotalCount: 2,
	}
)

func TestGetCodeDatasets(t *testing.T) {
	Convey("GetCodeDatasets returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Datasets{}, &expectedDatasets)
	})

	Convey("When datastore.GetCodeDatasets returns an error, then GetCodeDatasets returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/567/editions/890/codes/123/datasets", codeListURL), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return nil, ErrInternal
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("When datastore.GetCodeDatasets returns an invalid dataset (no ID), then GetCodeDatasets returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbmodels.Datasets{
					Items: []dbmodels.Dataset{{ID: ""}},
				}, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})

	Convey("Given a request to get code datasets", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		r.Header.Add("X-Forwarded-Proto", expectedProto)
		r.Header.Add("X-Forwarded-Host", expectedHost)
		r.Header.Add("X-Forwarded-Path-Prefix", expectedPathPrefix)

		Convey("When URL rewriting is disabled", func() {
			enableURLRewritingIsFalse := false
			api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewritingIsFalse)
			api.router.ServeHTTP(w, r)

			Convey("Then the response should have a status code of 200", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})

			Convey("And the response body should contain the original links", func() {
				var datasets models.Datasets
				err := json.Unmarshal(w.Body.Bytes(), &datasets)
				So(err, ShouldBeNil)

				So(datasets.Items[0].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s/datasets/%s", datasetURL, datasetID1))
				So(datasets.Items[1].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s/datasets/%s", datasetURL, datasetID2))

				So(datasets.Items[0].Editions[0].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s/datasets/%s/editions/%s", datasetURL, datasetID1, datasetEditionID))
				So(datasets.Items[0].Editions[0].Links.DatasetDimension.Href, ShouldEqual, fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d/dimensions/%s", datasetURL, datasetID1, datasetEditionID, latestDatasetEditionVersion, codeListID1))
				So(datasets.Items[0].Editions[0].Links.LatestVersion.Href, ShouldEqual, fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d", datasetURL, datasetID1, datasetEditionID, latestDatasetEditionVersion))

				So(datasets.Items[1].Editions, ShouldBeEmpty)
			})
		})

		Convey("When URL rewriting is enabled", func() {
			enableURLRewritingIsTrue := true
			api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewritingIsTrue)
			api.router.ServeHTTP(w, r)

			Convey("Then the response should have a status code of 200", func() {
				So(w.Code, ShouldEqual, http.StatusOK)
			})

			Convey("And the response body should contain the rewritten links", func() {
				var datasets models.Datasets
				err := json.Unmarshal(w.Body.Bytes(), &datasets)
				So(err, ShouldBeNil)

				So(datasets.Items[0].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s://%s/%s/datasets/%s", expectedProto, expectedHost, expectedPathPrefix, datasetID1))
				So(datasets.Items[1].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s://%s/%s/datasets/%s", expectedProto, expectedHost, expectedPathPrefix, datasetID2))

				So(datasets.Items[0].Editions[0].Links.Self.Href, ShouldEqual, fmt.Sprintf("%s://%s/%s/datasets/%s/editions/%s", expectedProto, expectedHost, expectedPathPrefix, datasetID1, datasetEditionID))
				So(datasets.Items[0].Editions[0].Links.DatasetDimension.Href, ShouldEqual, fmt.Sprintf("%s://%s/%s/datasets/%s/editions/%s/versions/%d/dimensions/%s", expectedProto, expectedHost, expectedPathPrefix, datasetID1, datasetEditionID, latestDatasetEditionVersion, codeListID1))
				So(datasets.Items[0].Editions[0].Links.LatestVersion.Href, ShouldEqual, fmt.Sprintf("%s://%s/%s/datasets/%s/editions/%s/versions/%d", expectedProto, expectedHost, expectedPathPrefix, datasetID1, datasetEditionID, latestDatasetEditionVersion))

				So(datasets.Items[1].Editions, ShouldBeEmpty)
			})
		})
	})
}

func TestGetCodeDatasets_Pagination(t *testing.T) {
	t.Parallel()

	Convey("When valid limit and offset query parameters are provided, then return code datasets information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?offset=0&limit=1", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Datasets{}, &datasetPaginationTestOne)
	})

	Convey("When valid limit above maximum and offset query parameters are provided, then return codes information according to the offset and limit", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?offset=1&limit=7", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Datasets{}, &datasetPaginationTestTwo)
	})

	Convey("When offset value greater than count provided, then return zero items", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?offset=3&limit=1", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		validateBody(w.Body, &models.Datasets{}, &datasetPaginationTestThree)
	})

	Convey("When a negative offset query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?offset=-1", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When a negative limit query parameter is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?limit=-1", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When limit above default maximum is provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?limit=1001", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})

	Convey("When non-integer query parameter value provided, then 400 status returned", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets?offset=x&limit=y", codeListURL, codeListID1, editionID1, codeID1), http.NoBody)
		w := httptest.NewRecorder()

		api := CreateCodeListAPI(mux.NewRouter(), &storetest.DataStoreMock{}, codeListAPIURL, datasetAPIURL, defaultOffset, defaultLimit, maxLimit, enableURLRewriting)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusBadRequest)
	})
}
