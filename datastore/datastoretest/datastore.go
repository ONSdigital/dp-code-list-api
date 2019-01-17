// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package storetest

import (
	"context"
	"github.com/ONSdigital/dp-code-list-api/models"
	"sync"
)

var (
	lockDataStoreMockGetCode         sync.RWMutex
	lockDataStoreMockGetCodeDatasets sync.RWMutex
	lockDataStoreMockGetCodeList     sync.RWMutex
	lockDataStoreMockGetCodeLists    sync.RWMutex
	lockDataStoreMockGetCodes        sync.RWMutex
	lockDataStoreMockGetEdition      sync.RWMutex
	lockDataStoreMockGetEditions     sync.RWMutex
)

// DataStoreMock is a mock implementation of DataStore.
//
//     func TestSomethingThatUsesDataStore(t *testing.T) {
//
//         // make and configure a mocked DataStore
//         mockedDataStore := &DataStoreMock{
//             GetCodeFunc: func(ctx context.Context, codeListID string, edition string, code string) (*models.Code, error) {
// 	               panic("TODO: mock out the GetCode method")
//             },
//             GetCodeDatasetsFunc: func(ctx context.Context, codeListID string, edition string, code string) (*models.Datasets, error) {
// 	               panic("TODO: mock out the GetCodeDatasets method")
//             },
//             GetCodeListFunc: func(ctx context.Context, codeListID string) (*models.CodeList, error) {
// 	               panic("TODO: mock out the GetCodeList method")
//             },
//             GetCodeListsFunc: func(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
// 	               panic("TODO: mock out the GetCodeLists method")
//             },
//             GetCodesFunc: func(ctx context.Context, codeListID string, edition string) (*models.CodeResults, error) {
// 	               panic("TODO: mock out the GetCodes method")
//             },
//             GetEditionFunc: func(ctx context.Context, codeListID string, edition string) (*models.Edition, error) {
// 	               panic("TODO: mock out the GetEdition method")
//             },
//             GetEditionsFunc: func(ctx context.Context, codeListID string) (*models.Editions, error) {
// 	               panic("TODO: mock out the GetEditions method")
//             },
//         }
//
//         // TODO: use mockedDataStore in code that requires DataStore
//         //       and then make assertions.
//
//     }
type DataStoreMock struct {
	// GetCodeFunc mocks the GetCode method.
	GetCodeFunc func(ctx context.Context, codeListID string, edition string, code string) (*models.Code, error)

	// GetCodeDatasetsFunc mocks the GetCodeDatasets method.
	GetCodeDatasetsFunc func(ctx context.Context, codeListID string, edition string, code string) (*models.Datasets, error)

	// GetCodeListFunc mocks the GetCodeList method.
	GetCodeListFunc func(ctx context.Context, codeListID string) (*models.CodeList, error)

	// GetCodeListsFunc mocks the GetCodeLists method.
	GetCodeListsFunc func(ctx context.Context, filterBy string) (*models.CodeListResults, error)

	// GetCodesFunc mocks the GetCodes method.
	GetCodesFunc func(ctx context.Context, codeListID string, edition string) (*models.CodeResults, error)

	// GetEditionFunc mocks the GetEdition method.
	GetEditionFunc func(ctx context.Context, codeListID string, edition string) (*models.Edition, error)

	// GetEditionsFunc mocks the GetEditions method.
	GetEditionsFunc func(ctx context.Context, codeListID string) (*models.Editions, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetCode holds details about calls to the GetCode method.
		GetCode []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
			// Edition is the edition argument value.
			Edition string
			// Code is the code argument value.
			Code string
		}
		// GetCodeDatasets holds details about calls to the GetCodeDatasets method.
		GetCodeDatasets []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
			// Edition is the edition argument value.
			Edition string
			// Code is the code argument value.
			Code string
		}
		// GetCodeList holds details about calls to the GetCodeList method.
		GetCodeList []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
		}
		// GetCodeLists holds details about calls to the GetCodeLists method.
		GetCodeLists []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// FilterBy is the filterBy argument value.
			FilterBy string
		}
		// GetCodes holds details about calls to the GetCodes method.
		GetCodes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
			// Edition is the edition argument value.
			Edition string
		}
		// GetEdition holds details about calls to the GetEdition method.
		GetEdition []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
			// Edition is the edition argument value.
			Edition string
		}
		// GetEditions holds details about calls to the GetEditions method.
		GetEditions []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CodeListID is the codeListID argument value.
			CodeListID string
		}
	}
}

// GetCode calls GetCodeFunc.
func (mock *DataStoreMock) GetCode(ctx context.Context, codeListID string, edition string, code string) (*models.Code, error) {
	if mock.GetCodeFunc == nil {
		panic("DataStoreMock.GetCodeFunc: method is nil but DataStore.GetCode was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
		Code       string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
		Edition:    edition,
		Code:       code,
	}
	lockDataStoreMockGetCode.Lock()
	mock.calls.GetCode = append(mock.calls.GetCode, callInfo)
	lockDataStoreMockGetCode.Unlock()
	return mock.GetCodeFunc(ctx, codeListID, edition, code)
}

// GetCodeCalls gets all the calls that were made to GetCode.
// Check the length with:
//     len(mockedDataStore.GetCodeCalls())
func (mock *DataStoreMock) GetCodeCalls() []struct {
	Ctx        context.Context
	CodeListID string
	Edition    string
	Code       string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
		Code       string
	}
	lockDataStoreMockGetCode.RLock()
	calls = mock.calls.GetCode
	lockDataStoreMockGetCode.RUnlock()
	return calls
}

// GetCodeDatasets calls GetCodeDatasetsFunc.
func (mock *DataStoreMock) GetCodeDatasets(ctx context.Context, codeListID string, edition string, code string) (*models.Datasets, error) {
	if mock.GetCodeDatasetsFunc == nil {
		panic("DataStoreMock.GetCodeDatasetsFunc: method is nil but DataStore.GetCodeDatasets was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
		Code       string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
		Edition:    edition,
		Code:       code,
	}
	lockDataStoreMockGetCodeDatasets.Lock()
	mock.calls.GetCodeDatasets = append(mock.calls.GetCodeDatasets, callInfo)
	lockDataStoreMockGetCodeDatasets.Unlock()
	return mock.GetCodeDatasetsFunc(ctx, codeListID, edition, code)
}

// GetCodeDatasetsCalls gets all the calls that were made to GetCodeDatasets.
// Check the length with:
//     len(mockedDataStore.GetCodeDatasetsCalls())
func (mock *DataStoreMock) GetCodeDatasetsCalls() []struct {
	Ctx        context.Context
	CodeListID string
	Edition    string
	Code       string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
		Code       string
	}
	lockDataStoreMockGetCodeDatasets.RLock()
	calls = mock.calls.GetCodeDatasets
	lockDataStoreMockGetCodeDatasets.RUnlock()
	return calls
}

// GetCodeList calls GetCodeListFunc.
func (mock *DataStoreMock) GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error) {
	if mock.GetCodeListFunc == nil {
		panic("DataStoreMock.GetCodeListFunc: method is nil but DataStore.GetCodeList was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
	}
	lockDataStoreMockGetCodeList.Lock()
	mock.calls.GetCodeList = append(mock.calls.GetCodeList, callInfo)
	lockDataStoreMockGetCodeList.Unlock()
	return mock.GetCodeListFunc(ctx, codeListID)
}

// GetCodeListCalls gets all the calls that were made to GetCodeList.
// Check the length with:
//     len(mockedDataStore.GetCodeListCalls())
func (mock *DataStoreMock) GetCodeListCalls() []struct {
	Ctx        context.Context
	CodeListID string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
	}
	lockDataStoreMockGetCodeList.RLock()
	calls = mock.calls.GetCodeList
	lockDataStoreMockGetCodeList.RUnlock()
	return calls
}

// GetCodeLists calls GetCodeListsFunc.
func (mock *DataStoreMock) GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
	if mock.GetCodeListsFunc == nil {
		panic("DataStoreMock.GetCodeListsFunc: method is nil but DataStore.GetCodeLists was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		FilterBy string
	}{
		Ctx:      ctx,
		FilterBy: filterBy,
	}
	lockDataStoreMockGetCodeLists.Lock()
	mock.calls.GetCodeLists = append(mock.calls.GetCodeLists, callInfo)
	lockDataStoreMockGetCodeLists.Unlock()
	return mock.GetCodeListsFunc(ctx, filterBy)
}

// GetCodeListsCalls gets all the calls that were made to GetCodeLists.
// Check the length with:
//     len(mockedDataStore.GetCodeListsCalls())
func (mock *DataStoreMock) GetCodeListsCalls() []struct {
	Ctx      context.Context
	FilterBy string
} {
	var calls []struct {
		Ctx      context.Context
		FilterBy string
	}
	lockDataStoreMockGetCodeLists.RLock()
	calls = mock.calls.GetCodeLists
	lockDataStoreMockGetCodeLists.RUnlock()
	return calls
}

// GetCodes calls GetCodesFunc.
func (mock *DataStoreMock) GetCodes(ctx context.Context, codeListID string, edition string) (*models.CodeResults, error) {
	if mock.GetCodesFunc == nil {
		panic("DataStoreMock.GetCodesFunc: method is nil but DataStore.GetCodes was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
		Edition:    edition,
	}
	lockDataStoreMockGetCodes.Lock()
	mock.calls.GetCodes = append(mock.calls.GetCodes, callInfo)
	lockDataStoreMockGetCodes.Unlock()
	return mock.GetCodesFunc(ctx, codeListID, edition)
}

// GetCodesCalls gets all the calls that were made to GetCodes.
// Check the length with:
//     len(mockedDataStore.GetCodesCalls())
func (mock *DataStoreMock) GetCodesCalls() []struct {
	Ctx        context.Context
	CodeListID string
	Edition    string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
	}
	lockDataStoreMockGetCodes.RLock()
	calls = mock.calls.GetCodes
	lockDataStoreMockGetCodes.RUnlock()
	return calls
}

// GetEdition calls GetEditionFunc.
func (mock *DataStoreMock) GetEdition(ctx context.Context, codeListID string, edition string) (*models.Edition, error) {
	if mock.GetEditionFunc == nil {
		panic("DataStoreMock.GetEditionFunc: method is nil but DataStore.GetEdition was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
		Edition:    edition,
	}
	lockDataStoreMockGetEdition.Lock()
	mock.calls.GetEdition = append(mock.calls.GetEdition, callInfo)
	lockDataStoreMockGetEdition.Unlock()
	return mock.GetEditionFunc(ctx, codeListID, edition)
}

// GetEditionCalls gets all the calls that were made to GetEdition.
// Check the length with:
//     len(mockedDataStore.GetEditionCalls())
func (mock *DataStoreMock) GetEditionCalls() []struct {
	Ctx        context.Context
	CodeListID string
	Edition    string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
		Edition    string
	}
	lockDataStoreMockGetEdition.RLock()
	calls = mock.calls.GetEdition
	lockDataStoreMockGetEdition.RUnlock()
	return calls
}

// GetEditions calls GetEditionsFunc.
func (mock *DataStoreMock) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
	if mock.GetEditionsFunc == nil {
		panic("DataStoreMock.GetEditionsFunc: method is nil but DataStore.GetEditions was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CodeListID string
	}{
		Ctx:        ctx,
		CodeListID: codeListID,
	}
	lockDataStoreMockGetEditions.Lock()
	mock.calls.GetEditions = append(mock.calls.GetEditions, callInfo)
	lockDataStoreMockGetEditions.Unlock()
	return mock.GetEditionsFunc(ctx, codeListID)
}

// GetEditionsCalls gets all the calls that were made to GetEditions.
// Check the length with:
//     len(mockedDataStore.GetEditionsCalls())
func (mock *DataStoreMock) GetEditionsCalls() []struct {
	Ctx        context.Context
	CodeListID string
} {
	var calls []struct {
		Ctx        context.Context
		CodeListID string
	}
	lockDataStoreMockGetEditions.RLock()
	calls = mock.calls.GetEditions
	lockDataStoreMockGetEditions.RUnlock()
	return calls
}
