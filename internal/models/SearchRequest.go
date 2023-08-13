package models

type SearchRequest struct {
	SearchKey string                 `json:"search_key"`
	Page      int                    `json:"page"`
	PageSize  int                    `json:"page_size"`
	SortBy    string                 `json:"sort_by"`
	SortOrder string                 `json:"sort_order"`
	Filters   map[string]interface{} `json:"filters"`
}
