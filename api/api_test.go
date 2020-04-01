package api

import "github.com/pkg/errors"

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

// ErrInternal is the error returned by mocks to emulate an internal service error
var ErrInternal = errors.New("internal error")
