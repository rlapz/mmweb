package api

type ApiResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}

type ApiRespPagination struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"total_items"`
	TotalPages int `json:"total_pages"`

	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

type ApiRespBodyList struct {
	List       any               `json:"list"`
	Pagination ApiRespPagination `json:"pagination"`
}
