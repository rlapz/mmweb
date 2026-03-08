package api

import "net/url"

type Response struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       any                 `json:"data,omitempty"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

type ResponsePagination struct {
	Page      int    `json:"page"`
	PageCap   int    `json:"page_cap"`
	ListLimit int    `json:"list_limit"`
	ListLen   int    `json:"list_len"`
	ListCap   int    `json:"list_cap"`
	Sort      string `json:"sort"`
	Order     string `json:"order"`

	PageNext string `json:"page_next"`
	PagePrev string `json:"page_prev"`
}

type RequestQuery struct {
	Id        string
	Page      int
	ListLimit int
	Sort      string
	Order     string
	Offset    int
	Query     url.Values
}
