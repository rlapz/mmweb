package api

import (
	"fmt"
	"net/url"

	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/util"
)

type ApiResp struct {
	Success    bool               `json:"success"`
	Message    string             `json:"message"`
	Data       any                `json:"data,omitempty"`
	Pagination *ApiRespPagination `json:"pagination,omitempty"`
}

type ApiRespPagination struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	ItemsLen   int    `json:"items_len"`
	ItemsTotal int    `json:"items_total"`
	PagesTotal int    `json:"pages_total"`
	Sort       string `json:"sort"`
	Order      string `json:"order"`

	Next string `json:"next,omitempty"`
	Prev string `json:"prev,omitempty"`
}

func (q *ApiRespPagination) BuildNavs(url string) {
	if q.Page < q.PagesTotal {
		q.Next = fmt.Sprintf("%s/?page=%d&limit=%d&sort=%s&order=%s", url, q.Page+1,
			q.Limit, q.Sort, q.Order)
	}

	if q.Page > 0 {
		q.Next = fmt.Sprintf("%s/?page=%d&limit=%d&sort=%s&order=%s", url, q.Page-1,
			q.Limit, q.Sort, q.Order)
	}
}

type ApiBaseQuery struct {
	Page   int
	Limit  int
	Offset int
	Sort   string
	Order  string
}

func (a *ApiBaseQuery) Parse(raw url.Values) {
	page := max(util.ParseInt(raw.Get("page")), 1)
	limit := max(util.ParseInt(raw.Get("limit")), config.DEF_RECORD_MIN_PER_PAGE_COUNT)
	limit = min(limit, config.DEF_RECORD_MAX_PER_PAGE_COUNT)

	a.Sort = raw.Get("sort")
	a.Order = raw.Get("order")
	a.Limit = limit
	a.Page = page
	a.Offset = (page - 1) * limit
}
