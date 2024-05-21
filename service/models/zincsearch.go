package models

type BulkV2Request struct {
	Index   string      `json:"index"`
	Records interface{} `json:"records"`
}

type SearchRequest struct {
	SearchType string      `json:"search_type"`
	Query      SearchQuery `json:"query"`
	SortFiels  []string    `json:"sort_fields"`
	From       int         `json:"from"`
	MaxResults int         `json:"max_results"`
}

type SearchQuery struct {
	Term string `json:"term"`
}

type SearchResponse struct {
	Hits struct {
		Hits []struct {
			ID        string                 `json:"_id"`
			Timestamp string                 `json:"@timestamp"`
			Score     float64                `json:"_score"`
			Source    map[string]interface{} `json:"_source"`
		} `json:"hits"`
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`
	} `json:"hits"`
	TimedOut bool    `json:"timed_out"`
	Took     float64 `json:"took"`
}
