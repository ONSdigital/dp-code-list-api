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

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

// Dataset models for testing
var (
	dbDataset = dbmodels.Dataset{
		ID:             codeID,
		DimensionLabel: "testDimensionLabel",
		Editions: []dbmodels.DatasetEdition{
			dbmodels.DatasetEdition{
				ID:            datasetEditionID,
				CodeListID:    codeID,
				LatestVersion: latestDatasetEditionVersion,
			},
		},
	}
	dbDatasets = dbmodels.Datasets{
		Items: []dbmodels.Dataset{dbDataset},
	}

	expectedDataset = models.Dataset{
		ID:             codeID,
		DimensionLabel: "testDimensionLabel",
		Editions: []models.DatasetEdition{
			models.DatasetEdition{
				ID:            datasetEditionID,
				LatestVersion: latestDatasetEditionVersion,
				Links: &models.DatasetEditionLinks{
					Self:             &models.Link{ID: datasetEditionID, Href: fmt.Sprintf("%s/datasets/%s/editions/%s", datasetURL, codeID, datasetEditionID)},
					DatasetDimension: &models.Link{ID: codeListID, Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d/dimensions/%s", datasetURL, codeID, datasetEditionID, latestDatasetEditionVersion, codeListID)},
					LatestVersion:    &models.Link{ID: fmt.Sprintf("%d", latestDatasetEditionVersion), Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/%d", datasetURL, codeID, datasetEditionID, latestDatasetEditionVersion)},
				},
			},
		},
		Links: &models.DatasetLinks{
			Self: &models.Link{ID: codeID, Href: fmt.Sprintf("%s/datasets/%s", datasetURL, codeID)},
		},
	}
	expectedDatasets = models.Datasets{
		Items:      []models.Dataset{expectedDataset},
		Count:      1,
		Offset:     0,
		Limit:      1,
		TotalCount: 1,
	}
)

func TestGetCodeDatasets(t *testing.T) {

	Convey("Get code lists returns a status of Ok", t, func() {
		r := httptest.NewRequest("GET", fmt.Sprintf("%s/code-lists/%s/editions/%s/codes/%s/datasets", codeListURL, codeListID, editionID, codeID), nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return &dbDatasets, nil
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusOK)

		payload, err := ioutil.ReadAll(w.Body)
		So(err, ShouldBeNil)
		apiDatasets := models.Datasets{}
		json.Unmarshal(payload, &apiDatasets)
		So(apiDatasets, ShouldResemble, expectedDatasets)
	})

	Convey("Get code lists returns a status of internal error", t, func() {
		r := httptest.NewRequest("GET", "http://localhost:8080/code-lists/567/editions/890/codes/123/datasets", nil)
		w := httptest.NewRecorder()

		mockDatastore := &storetest.DataStoreMock{
			GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*dbmodels.Datasets, error) {
				return nil, InternalError
			},
		}

		api := CreateCodeListAPI(mux.NewRouter(), mockDatastore, codeListURL, datasetURL)
		api.router.ServeHTTP(w, r)
		So(w.Code, ShouldEqual, http.StatusInternalServerError)
	})
}
