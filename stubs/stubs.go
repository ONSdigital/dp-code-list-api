package stubs

import "encoding/json"

// CodeList ...
type CodeList struct {
	ID    string               `json:"id"`
	Name  string               `json:"name"`
	Codes map[string]CodeEntry `json:"codes"`
}

// CodeEntry ...
type CodeEntry struct {
	ID    string `json:"id"`
	Code  string `json:"code"`
	Label string `json:"label"`
}

// GoodsAndServicesIndices ...
var GoodsAndServicesIndices CodeList

// Time ...
var Time CodeList

func init() {
	b, err := Asset("generators/output/goodsandservices.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &GoodsAndServicesIndices)

	b, err = Asset("generators/output/time.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &Time)
}
