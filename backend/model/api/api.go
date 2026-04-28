package api

import "net/url"

type Response struct {
	Success    bool                `json:"success"`
	Message    string              `json:"message"`
	Data       any                 `json:"data,omitempty"`
	Pagination *ResponsePagination `json:"pagination,omitempty"`
}

type ResponsePagination struct {
	Page    int `json:"page"`
	PageCap int `json:"page_cap"`
	Limit   int `json:"limit"`
	Len     int `json:"len"`
	Cap     int `json:"cap"`

	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

type RequestQuery struct {
	Id    string
	Page  int
	Limit int
	Query url.Values
}
