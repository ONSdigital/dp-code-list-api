package api

import "github.com/pkg/errors"

// InternalError is the error returned by mocks to emulate an internal service error
var InternalError = errors.New("internal error")

// Constants used by tests
const (
	codeListURL                 = "http://codelist-url"
	datasetURL                  = "http://dataset-url"
	codeListID                  = "testCodeList"
	codeID                      = "testCode"
	editionID                   = "testEdition"
	datasetEditionID            = "testDatasetEdition"
	latestDatasetEditionVersion = 248
)
