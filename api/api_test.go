package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

// Constants used by tests
const (
	codeListURL                 = "http://codelist-url"
	datasetURL                  = "http://dataset-url"
	codeListID                  = "testCodeList"
	codeID                      = "testCode"
	editionID                   = "testEdition"
	datasetEditionID            = "testDatasetEdition"
	latestDatasetEditionVersion = 248
	defaultOffset               = 0
	defaultLimit                = 100
)

// ErrInternal is the error returned by mocks to emulate an internal service error
var ErrInternal = errors.New("internal error")

// validateBody is a generic function that un-marshals the body to a provided interface, and validates it against the expected.
func validateBody(body *bytes.Buffer, apiStruct, expected interface{}) {
	payload, err := ioutil.ReadAll(body)
	So(err, ShouldBeNil)
	json.Unmarshal(payload, apiStruct)
	So(apiStruct, ShouldResemble, expected)
}
