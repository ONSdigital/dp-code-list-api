package models_test

import (
	"errors"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewEdition(t *testing.T) {
	Convey("NewEdition function called with a nil argument results in an empty Edition model", t, func() {
		edition := models.NewEdition(nil)
		So(edition, ShouldResemble, &models.Edition{})
	})

	Convey("Given a valid a valid edition database model", t, func() {
		dbEdition := &dbmodels.Edition{
			ID:    "testEdition",
			Label: "testLabel",
		}

		Convey("NewEdition function returns the corresponding API Edition model", func() {
			edition := models.NewEdition(dbEdition)
			So(edition, ShouldResemble, &models.Edition{
				ID:    "testEdition",
				Label: "testLabel",
			})
		})
	})
}

func TestNewEditions(t *testing.T) {
	Convey("NewEditions function called with a nil argument results in an empty API Editions model", t, func() {
		editions := models.NewEditions(nil)
		So(editions, ShouldResemble, &models.Editions{})
	})

	Convey("Given a valid editions Datasets model", t, func() {
		dbEditions := []dbmodels.Edition{
			{
				ID:    "testEdition",
				Label: "testLabel",
			},
		}

		Convey("NewEditions function returns the corresponding API Editions model", func() {
			editions := models.NewEditions(dbEditions)
			So(editions, ShouldResemble, &models.Editions{
				Items: []models.Edition{
					{
						ID:    "testEdition",
						Label: "testLabel",
					},
				},
			})
		})
	})
}

func TestEditionUpdateLinks(t *testing.T) {
	Convey("Given an Edition struct without ID", t, func() {
		edition := models.Edition{}

		Convey("UpdateLinks fails with the expected error", func() {
			err := edition.UpdateLinks("codelist1", "url1")
			So(err, ShouldResemble, errors.New("unable to create links - edition id not provided"))
		})
	})

	Convey("Given a valid Dataset struct with all Dataset items containing IDs", t, func() {
		edition := models.Edition{
			ID:    "testEdition",
			Label: "testLabel",
		}

		editionWithLinks := models.Edition{
			ID:    "testEdition",
			Label: "testLabel",
			Links: &models.EditionLinks{
				Self: &models.Link{
					ID:   "testEdition",
					Href: "url1/code-lists/codelist1/editions/testEdition",
				},
				Editions: &models.Link{
					Href: "url1/code-lists/codelist1/editions",
				},
				Codes: &models.Link{
					Href: "url1/code-lists/codelist1/editions/testEdition/codes",
				},
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := edition.UpdateLinks("codelist1", "url1")
			So(err, ShouldBeNil)
			So(edition, ShouldResemble, editionWithLinks)
		})
	})
}
