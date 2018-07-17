package store

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/store/mock"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"io"
	"testing"
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

func TestNeoDataStore_EditionExists_QueryNeoError(t *testing.T) {
	Convey("given conn.QueryNeo returns an error", t, func() {
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return nil, testErr
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

		Convey("then EditionExists should return an error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(exists, ShouldBeFalse)
			So(err.Error(), ShouldContainSubstring, "error executing neo4j query")

			q := fmt.Sprintf(countEditions, testCodeListID, testEdition)
			So(conn.QueryNeoCalls(), ShouldHaveLength, 1)
			So(conn.QueryNeoCalls()[0].Query, ShouldEqual, q)
			So(conn.CloseCalls(), ShouldHaveLength, 1)
		})
	})
}

func TestNeoDataStore_EditionExists_TooManyResults(t *testing.T) {
	Convey("given conn.QueryNeo returns more than 1 result", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				if index < 2 {
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

		Convey("then EditionExists should return an error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(exists, ShouldBeFalse)
			So(err.Error(), ShouldContainSubstring, "extract row result error: expected single result but was not")

			q := fmt.Sprintf(countEditions, testCodeListID, testEdition)
			So(conn.QueryNeoCalls(), ShouldHaveLength, 1)
			So(conn.QueryNeoCalls()[0].Query, ShouldEqual, q)
			So(conn.CloseCalls(), ShouldHaveLength, 1)

			So(rows.NextNeoCalls(), ShouldHaveLength, 2)
			So(rows.CloseCalls(), ShouldHaveLength, 1)
		})
	})
}

func TestNeoDataStore_EditionExists_ResultIncorrectType(t *testing.T) {
	Convey("given conn.QueryNeo does not return an int64 type", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				if index < 1 {
					index++
					return []interface{}{"this is not an int64"}, nil, nil
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

		Convey("then EditionExists should return an error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(exists, ShouldBeFalse)
			So(err.Error(), ShouldContainSubstring, "extract row result error: failed to cast result to int64")

			q := fmt.Sprintf(countEditions, testCodeListID, testEdition)
			So(conn.QueryNeoCalls(), ShouldHaveLength, 1)
			So(conn.QueryNeoCalls()[0].Query, ShouldEqual, q)
			So(conn.CloseCalls(), ShouldHaveLength, 1)

			So(rows.NextNeoCalls(), ShouldHaveLength, 1)
			So(rows.CloseCalls(), ShouldHaveLength, 1)
		})
	})
}

func TestNeoDataStore_EditionExists_NonUniqueEdition(t *testing.T) {
	Convey("given conn.QueryNeo returns count > 1", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				if index < 1 {
					index++
					return []interface{}{int64(2)}, nil, nil
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

		Convey("then EditionExists should return an error", func() {
			exists, err := neo.EditionExists(context.Background(), testCodeListID, "2018")
			So(exists, ShouldBeFalse)
			So(err.Error(), ShouldContainSubstring, "editionExists: multiple editions found")

			q := fmt.Sprintf(countEditions, testCodeListID, testEdition)
			So(conn.QueryNeoCalls(), ShouldHaveLength, 1)
			So(conn.QueryNeoCalls()[0].Query, ShouldEqual, q)
			So(conn.CloseCalls(), ShouldHaveLength, 1)

			So(rows.NextNeoCalls(), ShouldHaveLength, 2)
			So(rows.CloseCalls(), ShouldHaveLength, 1)
		})
	})
}
