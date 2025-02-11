package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/url"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	codeListAPIURL = &url.URL{Scheme: "http", Host: "localhost:22400"}
	datasetAPIURL  = &url.URL{Scheme: "http", Host: "localhost:22000"}
)

// Constants used by tests
const (
	codeListURL                 = "http://localhost:22400"
	datasetURL                  = "http://localhost:22000"
	codeListID1                 = "testCodeList1"
	codeListID2                 = "testCodeList2"
	codeID1                     = "testCode1"
	codeID2                     = "testCode2"
	datasetID1                  = "testDataset1"
	datasetID2                  = "testDataset2"
	editionID1                  = "testEdition1"
	editionID2                  = "testEdition2"
	datasetEditionID            = "testDatasetEdition"
	latestDatasetEditionVersion = 248
	defaultOffset               = 0
	defaultLimit                = 20
	maxLimit                    = 1000
	enableURLRewriting          = false
	expectedProto               = "https"
	expectedHost                = "api.example.com"
	expectedPathPrefix          = "v1"
)

// ErrInternal is the error returned by mocks to emulate an internal service error
var ErrInternal = errors.New("internal error")

// validateBody is a generic function that un-marshals the body to a provided interface, and validates it against the expected.
func validateBody(body *bytes.Buffer, apiStruct, expected interface{}) {
	payload, err := io.ReadAll(body)
	So(err, ShouldBeNil)
	json.Unmarshal(payload, apiStruct)
	So(apiStruct, ShouldResemble, expected)
}
