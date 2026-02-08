package util

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/rlapz/mmweb/model/api"
)

func httpResp(w http.ResponseWriter, code int, resp *api.ApiResp) {
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

	resp := api.ApiResp{
		Message: stb.String(),
	}

	httpResp(w, errCode, &resp)
	if err != nil {
		log.Printf("error: %s: %s\n", err.Error(), resp.Message)
	} else {
		log.Println("error:", resp.Message)
	}
}

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
	resp := api.ApiResp{
		Success: true,
		Message: msg,
		Data:    data,
	}

	httpResp(w, http.StatusOK, &resp)
}

func HttpCreated(w http.ResponseWriter, msg string, data any) {
	resp := api.ApiResp{
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

func HttpJsonParseBody[T any](reader io.Reader) (*T, error) {
	ret := new(T)
	if err := json.NewDecoder(reader).Decode(ret); err != nil {
		return nil, err
	}

	return ret, nil
}
