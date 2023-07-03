package models_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	testDatasetHost      = "datasetHost"
	testCodeListID       = "codeListID"
	testDatasetID        = "testDataset1"
	testDimensionLabel   = "myDimensionLabel"
	testDatasetEditionID = "testDatasetEdition1"
)

// DatasetEdition test variables
var (
	dbDatasetEdition1 = dbmodels.DatasetEdition{
		ID:            testDatasetEditionID,
		CodeListID:    testCodeListID,
		LatestVersion: 123,
	}

	datasetEdition1 = models.DatasetEdition{
		ID:            testDatasetEditionID,
		LatestVersion: 123,
	}

	datasetEdition1LinkSelf = models.Link{
		ID:   testDatasetEditionID,
		Href: fmt.Sprintf("%s/datasets/%s/editions/%s", testDatasetHost, testDatasetID, testDatasetEditionID),
	}

	datasetEdition1LinkLatestVersion = models.Link{
		ID:   "123",
		Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/123", testDatasetHost, testDatasetID, testDatasetEditionID),
	}

	datasetEdition1LinkDimension = models.Link{
		ID:   testCodeListID,
		Href: fmt.Sprintf("%s/datasets/%s/editions/%s/versions/123/dimensions/%s", testDatasetHost, testDatasetID, testDatasetEditionID, testCodeListID),
	}

	datasetEdition1WithLinks = models.DatasetEdition{
		ID:            testDatasetEditionID,
		LatestVersion: 123,
		Links: &models.DatasetEditionLinks{
			Self:             &datasetEdition1LinkSelf,
			DatasetDimension: &datasetEdition1LinkDimension,
			LatestVersion:    &datasetEdition1LinkLatestVersion,
		},
	}
)

// Dataset test variables
var (
	dbDataset1 = dbmodels.Dataset{
		ID:             testDatasetID,
		DimensionLabel: testDimensionLabel,
		Editions:       []dbmodels.DatasetEdition{dbDatasetEdition1},
	}

	dataset1 = models.Dataset{
		ID:             testDatasetID,
		DimensionLabel: testDimensionLabel,
		Editions:       []models.DatasetEdition{datasetEdition1},
	}

	dataset1WithLinks = models.Dataset{
		ID:             testDatasetID,
		DimensionLabel: testDimensionLabel,
		Links: &models.DatasetLinks{
			Self: &models.Link{
				ID:   testDatasetID,
				Href: fmt.Sprintf("%s/datasets/%s", testDatasetHost, testDatasetID),
			},
		},
		Editions: []models.DatasetEdition{
			datasetEdition1WithLinks,
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
			DimensionLabel: testDimensionLabel,
			Editions:       nil,
		}

		Convey("NewDataset function returns the corresponding API Dataset model", func() {
			dataset := models.NewDataset(dbDataset)
			So(dataset, ShouldResemble, &models.Dataset{
				ID:             "testID",
				DimensionLabel: testDimensionLabel,
				Editions:       []models.DatasetEdition{},
			})
		})
	})

	Convey("Given a valid database Dataset model", t, func() {
		Convey("NewDataset function returns the corresponding API Dataset model", func() {
			dataset := models.NewDataset(&dbDataset1)
			So(dataset, ShouldResemble, &dataset1)
		})
	})
}

func TestNewDatasets(t *testing.T) {
	Convey("NewDatasets function called with a nil argument results in an empty API Datasets model", t, func() {
		dataset := models.NewDatasets(nil)
		So(dataset, ShouldResemble, &models.Datasets{})
	})

	Convey("Given a valid database Datasets model", t, func() {
		dbDatasets := []dbmodels.Dataset{
			{
				ID: "id1",
			},
		}

		Convey("NewDatasets function returns the corresponding API Datasets model", func() {
			datasets := models.NewDatasets(dbDatasets)
			So(datasets, ShouldResemble, &models.Datasets{
				Items: []models.Dataset{
					{
						ID:       "id1",
						Editions: []models.DatasetEdition{},
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
		Convey("NewDatasetEdition function returns the corresponding API DatasetEdition model", func() {
			datasetEdition := models.NewDatasetEdition(&dbDatasetEdition1)
			So(datasetEdition, ShouldResemble, &datasetEdition1)
		})
	})
}

func TestDatasetsUpdateLinks(t *testing.T) {
	Convey("Given a Dataset struct that contains at least a Dataset item without ID", t, func() {
		datasets := models.Datasets{
			Items: []models.Dataset{
				{},
			},
		}

		Convey("UpdateLinks fails with the expected error, due to nested error", func() {
			err := datasets.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldResemble, errors.New("Error(s) happened updating dataset links"))
		})
	})

	Convey("Given a valid Dataset struct with all Dataset items containing IDs, but no editions", t, func() {
		datasets := models.Datasets{
			Items: []models.Dataset{
				{
					ID: testDatasetID,
				},
			},
		}

		datasetsWithLinks := models.Datasets{
			Items: []models.Dataset{
				{
					ID: testDatasetID,
					Links: &models.DatasetLinks{
						Self: &models.Link{
							ID:   testDatasetID,
							Href: fmt.Sprintf("%s/datasets/%s", testDatasetHost, testDatasetID),
						},
					},
				},
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := datasets.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldBeNil)
			So(datasets, ShouldResemble, datasetsWithLinks)
		})
	})

	Convey("Given a valid Dataset struct with all Dataset items containing IDs, with nested Editions", t, func() {
		datasets := models.Datasets{
			Items: []models.Dataset{
				dataset1,
			},
		}

		datasetsWithLinks := models.Datasets{
			Items: []models.Dataset{
				dataset1WithLinks,
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := datasets.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldBeNil)
			So(datasets, ShouldResemble, datasetsWithLinks)
		})
	})
}

func TestDatasetUpdateLinks(t *testing.T) {
	Convey("Given a Dataset struct without ID", t, func() {
		dataset := models.Dataset{}

		Convey("UpdateLinks fails with the expected error and dataset is not mutated", func() {
			err := dataset.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldResemble, errors.New("invalid dataset provided"))
			So(dataset, ShouldResemble, dataset)
		})
	})

	Convey("Given a valid Dataset with ID", t, func() {
		dataset := models.Dataset{
			ID: testDatasetID,
		}

		datasetWithLinks := models.Dataset{
			ID: testDatasetID,
			Links: &models.DatasetLinks{
				Self: &models.Link{
					ID:   testDatasetID,
					Href: fmt.Sprintf("%s/datasets/%s", testDatasetHost, testDatasetID),
				},
			},
		}

		Convey("UpdateLinks generates the expected link", func() {
			err := dataset.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldBeNil)
			So(dataset, ShouldResemble, datasetWithLinks)
		})
	})

	Convey("Given a valid Dataset with ID, and nested Edition with no ID", t, func() {
		dataset := models.Dataset{
			ID:             testDatasetID,
			DimensionLabel: testDimensionLabel,
			Editions: []models.DatasetEdition{
				{},
			},
		}

		datasetWithLinks := models.Dataset{
			ID:             testDatasetID,
			DimensionLabel: testDimensionLabel,
			Editions: []models.DatasetEdition{
				{},
			},
			Links: &models.DatasetLinks{
				Self: &models.Link{
					ID:   testDatasetID,
					Href: fmt.Sprintf("%s/datasets/%s", testDatasetHost, testDatasetID),
				},
			},
		}

		Convey("UpdateLinks generates the expected link, but fails due to the nested exception", func() {
			err := dataset.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldResemble, errors.New("Error(s) happened updating dataset edition links"))
			So(dataset, ShouldResemble, datasetWithLinks)
		})
	})

	Convey("Given a valid Dataset struct with all Dataset items containing IDs, with nested Editions", t, func() {
		Convey("UpdateLinks generates the expected links", func() {
			err := dataset1.UpdateLinks(testDatasetHost, testCodeListID)
			So(err, ShouldBeNil)
			So(dataset1, ShouldResemble, dataset1WithLinks)
		})
	})
}

func TestDatasetEditionUpdateLinks(t *testing.T) {
	Convey("Given a DatasetEdition struct without ID", t, func() {
		datasetEdition := models.DatasetEdition{
			LatestVersion: 246,
		}

		Convey("UpdateLinks fails with the expected error", func() {
			err := datasetEdition.UpdateLinks(testDatasetHost, testCodeListID, testDatasetID)
			So(err, ShouldResemble, errors.New("datasetEdition with empty ID"))
			So(datasetEdition.Links, ShouldResemble, &models.DatasetEditionLinks{
				Self:             nil,
				LatestVersion:    nil,
				DatasetDimension: nil,
			})
		})
	})

	Convey("Given a DatasetEdition struct without a valid LatestVersion", t, func() {
		datasetEdition := models.DatasetEdition{
			ID:            testDatasetEditionID,
			LatestVersion: -1,
		}

		Convey("UpdateLinks fails with the expected error", func() {
			err := datasetEdition.UpdateLinks(testDatasetHost, testCodeListID, testDatasetID)
			So(err, ShouldResemble, errors.New("datasetEdition with invalid LatestVersion"))
			So(datasetEdition.Links, ShouldResemble, &models.DatasetEditionLinks{
				Self:             &datasetEdition1LinkSelf,
				LatestVersion:    nil,
				DatasetDimension: nil,
			})
		})
	})

	Convey("Given a valid DatasetEdition", t, func() {
		Convey("UpdateLinks called with an empty codeListID fails with the expected error", func() {
			err := datasetEdition1.UpdateLinks(testDatasetHost, "", testDatasetID)
			So(err, ShouldResemble, errors.New("empty codeListID provided"))
			So(datasetEdition1.Links, ShouldResemble, &models.DatasetEditionLinks{
				Self:             &datasetEdition1LinkSelf,
				LatestVersion:    &datasetEdition1LinkLatestVersion,
				DatasetDimension: nil,
			})
		})

		Convey("UpdateLinks called with an empty datasetID fails with the expected error", func() {
			err := datasetEdition1.UpdateLinks(testDatasetHost, testCodeListID, "")
			So(err, ShouldResemble, errors.New("empty datasetID provided"))
			So(datasetEdition1.Links, ShouldResemble, &models.DatasetEditionLinks{
				Self:             &datasetEdition1LinkSelf,
				LatestVersion:    &datasetEdition1LinkLatestVersion,
				DatasetDimension: nil,
			})
		})

		Convey("UpdateLinks called with valid parameters results in the links being updated", func() {
			err := datasetEdition1.UpdateLinks(testDatasetHost, testCodeListID, testDatasetID)
			So(err, ShouldBeNil)
			So(datasetEdition1, ShouldResemble, datasetEdition1WithLinks)
			So(datasetEdition1.Links, ShouldResemble, &models.DatasetEditionLinks{
				Self:             &datasetEdition1LinkSelf,
				LatestVersion:    &datasetEdition1LinkLatestVersion,
				DatasetDimension: &datasetEdition1LinkDimension,
			})
		})
	})
}
