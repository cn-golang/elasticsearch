package es_search

import "encoding/json"

// 最外层
type ReqSearchData struct {
	Hits HitsData `json:"hits"`
}

// 第一层 hits
type HitsData struct {
	Total TotalData     `json:"total"`
	Hits  []HitsTwoData `json:"hits"`
}

// 第二层 hits
type HitsTwoData struct {
	Source json.RawMessage `json:"_source"`
}

// 第一层 total
type TotalData struct {
	Value    int
	Relation string
}
