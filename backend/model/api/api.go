package api

type ApiResp struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       any                `json:"data,omitempty"`
	Pagination *ApiRespPagination `json:"pagination,omitempty"`
}

type ApiRespPagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`

	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}
