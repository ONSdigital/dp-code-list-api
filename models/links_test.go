package models_test

import (
	"testing"

	"github.com/ONSdigital/dp-code-list-api/models"
	. "github.com/smartystreets/goconvey/convey"
)

var id = "id"

func TestCreateLink(t *testing.T) {
	Convey("Given a valid domain config and fully qualified url", t, func() {
		domain := "http://api.example.com/v1"
		url := "http://localhost:22400/code-list/local-authority/codes/E01000064"

		Convey("When the CreateLink function is called", func() {
			link := models.CreateLink(id, url, domain)

			Convey("Then the returned values should be as expected", func() {
				So(link.Href, ShouldEqual, "http://api.example.com/v1/code-list/local-authority/codes/E01000064")
				So(link.ID, ShouldEqual, id)
			})
		})
	})

	Convey("Given default config and fully qualified url", t, func() {
		domain := "http://localhost:22400"
		url := "http://somedomain/code-list/local-authority/codes/E01000064"

		Convey("When the CreateLink function is called", func() {
			link := models.CreateLink(id, url, domain)

			Convey("Then the returned values should be as expected", func() {
				So(link.Href, ShouldEqual, "http://localhost:22400/code-list/local-authority/codes/E01000064")
				So(link.ID, ShouldEqual, id)
			})
		})
	})

	Convey("Given default config and a path", t, func() {
		domain := "http://localhost:22400"
		url := "/code-list/local-authority/codes/E01000064"

		Convey("When the CreateLink function is called", func() {
			link := models.CreateLink(id, url, domain)

			Convey("Then the returned values should be as expected", func() {
				So(link.Href, ShouldEqual, "http://localhost:22400/code-list/local-authority/codes/E01000064")
				So(link.ID, ShouldEqual, id)
			})
		})
	})

	Convey("Given a malformed url", t, func() {
		domain := "http://localhost:22400<<"
		url := "/code-list/local-authority/codes/E01000064"

		Convey("When the CreateLink function is called", func() {
			link := models.CreateLink(id, url, domain)

			Convey("Then the link is not created, and a nil value is returned", func() {
				So(link, ShouldBeNil)
			})
		})
	})

	Convey("Given a malformed domain", t, func() {
		domain := "http://localhost:22400"
		url := "/code-list/local-authority/codes/=*-+!@£$%^&*()_"

		Convey("When the CreateLink function is called", func() {
			link := models.CreateLink(id, url, domain)

			Convey("Then the link is not created, and a nil value is returned", func() {
				So(link, ShouldBeNil)
			})
		})
	})
}
