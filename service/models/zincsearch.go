package models

type BulkV2Request struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}
