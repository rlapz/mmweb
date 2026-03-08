package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

/*
 * Http
 */
func HttpErrInternal(w http.ResponseWriter, err error, msg string) {
	httpRespErr(w, err, http.StatusInternalServerError, "internal", msg)
}

func HttpErrNotFound(w http.ResponseWriter, msg string) {
	httpRespErr(w, nil, http.StatusNotFound, "not found", msg)
}

func HttpErrBadRequest(w http.ResponseWriter, msg string) {
	httpRespErr(w, nil, http.StatusBadRequest, "bad request", msg)
}

func HttpErrUnauthorized(w http.ResponseWriter, msg string) {
	httpRespErr(w, nil, http.StatusUnauthorized, "unauthorized", msg)
}

func HttpOk(w http.ResponseWriter, msg string, data any) {
	resp := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}

	httpResp(w, http.StatusOK, &resp)
}

func HttpOkWithPagination(w http.ResponseWriter, msg string, data any, pag *ResponsePagination) {
	resp := Response{
		Success:    true,
		Message:    msg,
		Data:       data,
		Pagination: pag,
	}

	httpResp(w, http.StatusOK, &resp)
}

func HttpCreated(w http.ResponseWriter, msg string, data any) {
	resp := Response{
		Success: true,
		Message: msg,
		Data:    data,
	}

	httpResp(w, http.StatusCreated, &resp)
}

func HttpMethodCheck(w http.ResponseWriter, r *http.Request, expected string) bool {
	if r.Method != expected {
		HttpErrBadRequest(w, "invalid method")
		return false
	}

	return true
}

/*
 * ResponsePagination
 */
func ResponsePaginationNew(query *RequestQuery, pag *model.Pagination) *ResponsePagination {
	p := new(ResponsePagination)
	p.Page = pag.Page
	p.ListCap = pag.PageCap
	p.ListLimit = pag.ListLimit
	p.ListLen = pag.ListLen
	p.ListCap = pag.ListCap
	p.Sort = p.Order
	if query.Path == "" {
		return p
	}

	q := util.Cnd((query.Path == ""), "?", fmt.Sprintf("?%s&", query.Path))
	if p.Page < p.PageCap {
		p.PageNext = fmt.Sprintf("%s%spage=%d&list_limit=%d&sort=%s&order=%s",
			query.Path, q, p.Page+1, p.ListLimit, p.Sort, p.Order)
	}

	if p.Page > 0 {
		p.PagePrev = fmt.Sprintf("%s%spage=%d&list_limit=%d&sort=%s&order=%s",
			query.Path, q, p.Page-1, p.ListLimit, p.Sort, p.Order)
	}

	return p
}

/*
 * RequestQuery
 */
func RequestQueryParse(r *http.Request) *RequestQuery {
	raw := r.URL.Query()
	page := max(util.ParseInt(raw.Get("page")), 1)
	limit := max(util.ParseInt(raw.Get("limit")), config.DEF_RECORD_MIN_PER_PAGE_COUNT)
	limit = min(limit, config.DEF_RECORD_MAX_PER_PAGE_COUNT)

	return &RequestQuery{
		Id:        raw.Get("id"),
		Sort:      raw.Get("sort"),
		Order:     raw.Get("order"),
		ListLimit: limit,
		Page:      page,
		Offset:    (page - 1) * limit,
		Path:      r.URL.Path,
		Query:     raw,
	}
}

func (q *RequestQuery) Get(key string) string {
	return q.Query.Get(key)
}

/**************
 * Private    *
 **************/
func httpResp(w http.ResponseWriter, code int, resp *Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(&resp)
	if err != nil {
		log.Println("error: baseResp:", err.Error())
	}
}

func httpRespErr(w http.ResponseWriter, err error, errCode int, def string, msg string) {
	var stb strings.Builder
	stb.Grow(len(def) + len(msg) + 2)

	stb.WriteString(def)
	if msg != "" {
		fmt.Fprint(&stb, ": ", msg)
	}

	resp := Response{
		Message: stb.String(),
	}

	httpResp(w, errCode, &resp)
	if err != nil {
		log.Printf("error: %s: %s\n", err.Error(), resp.Message)
	} else {
		log.Println("error:", resp.Message)
	}
}
