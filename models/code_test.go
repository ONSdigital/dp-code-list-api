package models_test

import (
	"errors"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	dbmodels "github.com/ONSdigital/dp-graph/v2/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewCode(t *testing.T) {
	Convey("NewCode function called with a nil argument results in an empty API Code model", t, func() {
		code := models.NewCode(nil)
		So(code, ShouldResemble, &models.Code{})
	})

	Convey("Given a valid database Code model", t, func() {
		dbCode := &dbmodels.Code{
			ID:    "testID",
			Label: "testLabel",
			Code:  "testCode",
		}

		Convey("NewCode function returns the corresponding API Code model, with ID obtained from DB Code field", func() {
			code := models.NewCode(dbCode)
			So(code, ShouldResemble, &models.Code{
				ID:    "testCode",
				Label: "testLabel",
			})
		})
	})
}

func TestNewCodeResults(t *testing.T) {
	Convey("NewCodeResults function called with a nil argument results in an empty API CodeResults model", t, func() {
		codeResults := models.NewCodeResults(nil)
		So(codeResults, ShouldResemble, &models.CodeResults{})
	})

	Convey("Given a valid database CodeResults model with items", t, func() {
		dbCodeResults := []dbmodels.Code{
			{
				ID:    "testID",
				Label: "testLabel",
				Code:  "testCode",
			},
		}

		Convey("NewCodeResults function returns the corresponding API Code model", func() {
			codeResults := models.NewCodeResults(dbCodeResults)
			So(codeResults, ShouldResemble, &models.CodeResults{
				Items: []models.Code{
					{
						ID:    "testCode",
						Label: "testLabel",
					},
				},
			})
		})
	})
}

func TestCodeUpdateLinks(t *testing.T) {
	Convey("Given a Code struct without ID", t, func() {
		code := models.Code{}

		Convey("UpdateLinks fails with the expected error", func() {
			err := code.UpdateLinks("host1", "codelist1", "edition1")
			So(err, ShouldResemble, errors.New("unable to create links - code ID not provided"))
		})
	})

	Convey("Given a valid Code struct", t, func() {
		code := models.Code{
			ID:    "testCode",
			Label: "testLabel",
		}

		expectedCodeWithLinks := models.Code{
			ID:    "testCode",
			Label: "testLabel",
			Links: &models.CodeLinks{
				Self: &models.Link{
					ID:   "testCode",
					Href: "host1/code-lists/codelist1/editions/edition1/codes/testCode",
				},
				CodeList: &models.Link{
					Href: "host1/code-lists/codelist1",
				},
				Datasets: &models.Link{
					Href: "host1/code-lists/codelist1/editions/edition1/codes/testCode/datasets",
				},
			},
		}

		Convey("UpdateLinks generates the expected links using the Code value", func() {
			err := code.UpdateLinks("host1", "codelist1", "edition1")
			So(err, ShouldBeNil)
			So(code, ShouldResemble, expectedCodeWithLinks)
		})
	})
}
