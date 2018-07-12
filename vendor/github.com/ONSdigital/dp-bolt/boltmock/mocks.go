package boltmock

import "github.com/ONSdigital/dp-bolt/bolt"

type QueryParams struct {
	Query  string
	Params map[string]interface{}
}

type QueryFunc func(query string, params map[string]interface{}, resultExtractor bolt.ResultExtractor) error

type DB struct {
	QueryForResultCalls  []QueryParams
	QueryForResultFuncs  []QueryFunc
	CloseFunc            func() error
	QueryForResultsCalls  []QueryParams
	QueryForResultsFuncs  []QueryFunc
}

func (m *DB) QueryForResult(query string, params map[string]interface{}, resultExtractor bolt.ResultExtractor) error {
	if m.QueryForResultCalls == nil {
		m.QueryForResultCalls = []QueryParams{}
	}

	index := len(m.QueryForResultCalls)
	m.QueryForResultCalls = append(m.QueryForResultCalls, newQueryParams(query, params))
	return m.QueryForResultFuncs[index](query, params, resultExtractor)
}

func (m *DB) QueryForResults(query string, params map[string]interface{}, resultExtractor bolt.ResultExtractor) error {
	if m.QueryForResultsCalls == nil {
		m.QueryForResultsCalls = []QueryParams{}
	}

	index := len(m.QueryForResultsCalls)
	m.QueryForResultsCalls = append(m.QueryForResultsCalls, newQueryParams(query, params))
	return m.QueryForResultsFuncs[index](query, params, resultExtractor)
}

func (m *DB) Close() error {
	return m.CloseFunc()
}

func newQueryParams(query string, params map[string]interface{}) QueryParams {
	var p map[string]interface{}
	if params != nil {
		p = make(map[string]interface{})
		for k, v := range params {
			p[k] = v
		}
	}
	return QueryParams{Query: query, Params: p}
}
