package models_test

import (
	"errors"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	dbDataset1 = dbmodels.Dataset{
		ID:             "testID",
		DimensionLabel: "testDimensionLabel",
		Editions: []dbmodels.DatasetEdition{
			dbmodels.DatasetEdition{
				ID:            "testID",
				CodeListID:    "testCodeListID",
				LatestVersion: 123,
			},
		},
	}
)

func TestNewDataset(t *testing.T) {

	Convey("NewDataset function called with a nil argument results in an empty API Dataset model", t, func() {
		dataset := models.NewDataset(nil)
		So(dataset, ShouldResemble, &models.Dataset{})
	})

	Convey("Given a valid database Dataset model without editions", t, func() {

		dbDataset := &dbmodels.Dataset{
			ID:             "testID",
			DimensionLabel: "testDimensionLabel",
			Editions:       nil,
		}

		Convey("NewDataset function returns the corresponding API Dataset model", func() {
			datast := models.NewDataset(dbDataset)
			So(datast, ShouldResemble, &models.Dataset{
				ID:             "testID",
				DimensionLabel: "testDimensionLabel",
				Editions:       []models.DatasetEdition{},
			})
		})
	})

	Convey("Given a valid database Dataset model", t, func() {

		Convey("NewDataset function returns the corresponding API Dataset model", func() {
			datast := models.NewDataset(&dbDataset1)
			So(datast, ShouldResemble, &models.Dataset{
				ID:             "testID",
				DimensionLabel: "testDimensionLabel",
				Editions: []models.DatasetEdition{
					models.DatasetEdition{},
				},
			})
		})
	})
}

func TestNewDatasets(t *testing.T) {

	Convey("NewDatasets function called with a nil argument results in an empty API Datasets model", t, func() {
		dataset := models.NewDatasets(nil)
		So(dataset, ShouldResemble, &models.Datasets{})
	})

	Convey("Given a valid database Datasets model", t, func() {
		dbDatasets := &dbmodels.Datasets{
			Items: []dbmodels.Dataset{
				dbDataset1,
			},
		}

		Convey("NewDatasets function returns the corresponding API Datasets model", func() {
			datasets := models.NewDatasets(dbDatasets)
			So(datasets, ShouldResemble, &models.Datasets{
				Items: []models.Dataset{
					models.Dataset{
						ID:             "testID",
						DimensionLabel: "testDimensionLabel",
						Editions: []models.DatasetEdition{
							models.DatasetEdition{},
						},
					},
				},
			})
		})
	})
}

func TestNewDatasetEdition(t *testing.T) {

	Convey("NewDatasetEdition function called with a nil argument results in an empty API Dataset model", t, func() {
		datasetEdition := models.NewDatasetEdition(nil)
		So(datasetEdition, ShouldResemble, &models.DatasetEdition{})
	})

	Convey("Given a valid database DatasetEdition model", t, func() {

		dbDatasetEdition := &dbmodels.DatasetEdition{
			ID:            "testID",
			CodeListID:    "testCodeList",
			LatestVersion: 123,
		}

		Convey("NewDatasetEdition function returns the corresponding API DatasetEdition model", func() {
			datasetEdition := models.NewDatasetEdition(dbDatasetEdition)
			So(datasetEdition, ShouldResemble, &models.DatasetEdition{})
		})
	})
}

func TestDatasetsUpdateLinks(t *testing.T) {

	Convey("Given a Dataset struct that contains at least a Dataset item without ID", t, func() {

		datasets := models.Datasets{
			Items: []models.Dataset{
				models.Dataset{},
			},
		}

		Convey("UpdateLinks fails with the expected error", func() {
			err := datasets.UpdateLinks("host1", "dataset1", "codelist1", "edition1", "code1")
			So(err, ShouldResemble, errors.New("invalid dataset provided"))
		})
	})

	Convey("Given a valid Dataset struct with all Dataset items containing IDs", t, func() {

		datasets := models.Datasets{
			Items: []models.Dataset{
				models.Dataset{
					ID: "testID",
				},
			},
		}

		datasetsWithLinks := models.Datasets{
			Items: []models.Dataset{
				models.Dataset{
					ID: "testID",
					Links: &models.DatasetLinks{
						Self: &models.Link{
							ID:   "testID",
							Href: "dataset1/datasets/testID",
						},
					},
				},
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := datasets.UpdateLinks("host1", "dataset1", "codelist1", "edition1", "code1")
			So(err, ShouldBeNil)
			So(datasets, ShouldResemble, datasetsWithLinks)
		})
	})

	// TODO test with nested editions

}
