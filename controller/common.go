package controller

import (
	"fmt"

	"github.com/rlapz/mmweb/model/api"
)

func PaginationSetNav(baseUrl string, pag *api.ApiRespPagination) *api.ApiRespPagination {
	if pag.Page < pag.TotalPages {
		pag.Next = fmt.Sprintf("%s?page=%d", baseUrl, pag.Page+1)
	}

	if pag.Page > 0 {
		pag.Prev = fmt.Sprintf("%s?page=%d", baseUrl, pag.Page-1)
	}

	return pag
}
