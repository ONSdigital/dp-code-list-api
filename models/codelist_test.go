package models_test

import (
	"errors"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCodeList(t *testing.T) {
	Convey("NewCodeList function called with a nil argument results in an empty API CodeList model", t, func() {
		codeList := models.NewCodeList(nil)
		So(codeList, ShouldResemble, &models.CodeList{})
	})

	Convey("Given a valid database CodeList model", t, func() {
		dbCodeList := &dbmodels.CodeList{
			ID: "testID",
		}

		Convey("NewCodeList function returns the corresponding API CodeList model", func() {
			codeList := models.NewCodeList(dbCodeList)
			So(codeList, ShouldResemble, &models.CodeList{
				ID: "testID",
			})
		})
	})
}

func TestNewCodeListResults(t *testing.T) {
	Convey("NewCodeListResults function called with a nil argument results in an empty API CodeListResults model", t, func() {
		codeListResults := models.NewCodeListResults(nil)
		So(codeListResults, ShouldResemble, &models.CodeListResults{})
	})

	Convey("Given a valid database CodeListResults model with items", t, func() {
		dbCodeListResults := []dbmodels.CodeList{
			{
				ID: "testID",
			},
		}

		Convey("NewCodeListResults function returns the corresponding API CodeList model", func() {
			codeListResults := models.NewCodeListResults(dbCodeListResults)
			So(codeListResults, ShouldResemble, &models.CodeListResults{
				Items: []models.CodeList{
					{
						ID: "testID",
					},
				},
			})
		})
	})
}

func TestCodeListUpdateLinks(t *testing.T) {
	Convey("Given a CodeList struct without ID", t, func() {
		codeList := models.CodeList{}

		Convey("UpdateLinks fails with the expected error", func() {
			err := codeList.UpdateLinks("url1")
			So(err, ShouldResemble, errors.New("unable to create links - codelist id not provided"))
		})
	})

	Convey("Given a valid CodeList struct", t, func() {
		codeList := models.CodeList{
			ID: "codelistID",
		}

		codeListWithLinks := models.CodeList{
			ID: "codelistID",
			Links: &models.CodeListLink{
				Self: &models.Link{
					ID:   "codelistID",
					Href: "testURL/code-lists/codelistID",
				},
				Editions: &models.Link{
					Href: "testURL/code-lists/codelistID/editions",
				},
			},
		}

		Convey("UpdateLinks generates the expected links", func() {
			err := codeList.UpdateLinks("testURL")
			So(err, ShouldBeNil)
			So(codeList, ShouldResemble, codeListWithLinks)
		})
	})
}
