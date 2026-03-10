package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/model"
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
	return &ResponsePagination{
		Page:      pag.Page,
		PageCap:   pag.PageCap,
		ListLimit: pag.ListLimit,
		ListLen:   pag.ListLen,
		ListCap:   pag.ListCap,
		Sort:      pag.Order,
	}
}

/*
 * RequestQuery
 */
func RequestQueryParse(r *http.Request) *RequestQuery {
	raw := r.URL.Query()
	page := max(parseInt(raw.Get("page")), 1)
	limit := max(parseInt(raw.Get("limit")), config.DEF_RECORD_MIN_PER_PAGE_COUNT)
	limit = min(limit, config.DEF_RECORD_MAX_PER_PAGE_COUNT)

	return &RequestQuery{
		Id:        raw.Get("id"),
		Sort:      raw.Get("sort"),
		Order:     raw.Get("order"),
		ListLimit: limit,
		Page:      page,
		Offset:    (page - 1) * limit,
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

func parseInt(str string) int {
	val, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return -1
	}

	return int(val)
}
