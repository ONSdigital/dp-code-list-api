package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-code-list-api/store/mock"
	"context"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"io"
	"fmt"
	"github.com/pkg/errors"
)

var (
	closeNoErr = func() error {
		return nil
	}

	testCodeListID = "666" // number of the beast \m/
	testEdition    = "2018"
	testErr        = errors.New("error")
)

func TestNeoDataStore_EditionExists(t *testing.T) {
	Convey("given neo4j returns count == 1 for specified codelist & edition", t, func() {

		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				if index < 1 {
					index++
					return []interface{}{int64(1)}, nil, nil
				}
				return nil, nil, io.EOF
			},
			CloseFunc: closeNoErr,
		}

		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return rows, nil
			},
		}

		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return conn, nil
			},
		}

		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}

		Convey("then EditionExists should return true and no error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(err, ShouldBeNil)
			So(exists, ShouldBeTrue)
			So(pool.OpenPoolCalls(), ShouldHaveLength, 1)

			q := fmt.Sprintf(countEditions, testCodeListID, testEdition)
			So(conn.QueryNeoCalls(), ShouldHaveLength, 1)
			So(conn.QueryNeoCalls()[0].Query, ShouldEqual, q)
			So(conn.CloseCalls(), ShouldHaveLength, 1)

			So(rows.NextNeoCalls(), ShouldHaveLength, 2)
			So(rows.CloseCalls(), ShouldHaveLength, 1)
		})
	})
}

func TestNeoDataStore_EditionExists_OpenConnError(t *testing.T) {
	Convey("given pool.OpenPool returns an error", t, func() {
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return nil, testErr
			},
		}

		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}

		Convey("then EditionExists should return an error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(exists, ShouldBeFalse)
			So(err.Error(), ShouldContainSubstring, "error while attempting to open neo4j connection")
		})
	})
}
