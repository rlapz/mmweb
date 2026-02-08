package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/model/api"
	"github.com/rlapz/mmweb/util"
)

func (c *controller) todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getTodoList(w, r)
	case http.MethodPost:
		c.postTodoItem(w, r)
	case http.MethodPut:
		c.putTodoItem(w, r)
	case http.MethodDelete:
		c.deleteTodoItem(w, r)
	default:
		util.HttpMethodCheck(w, r, "invalid")
	}
}

func (c *controller) getTodoList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	//id := query.Get("id")
	// if id == "" -> show all todos, else -> show detail todo

	claims := util.ContextGetJwtClaims(r.Context())

	body := api.ApiRespBodyList{
		List: []string{
			claims["iss"].(string),
			claims["jti"].(string),
			query.Get("id"),
			fmt.Sprint(query),
		},
	}

	util.HttpOk(w, "ok", body)
}

func (c *controller) postTodoItem(w http.ResponseWriter, r *http.Request) {
	todo, err := util.HttpJsonParseBody[model.Todo](r.Body)
	if err != nil {
		util.HttpErrBadRequest(w, "")
		return
	}

	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	uname, err := claims.GetIssuer()
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get context claims")
		return
	}

	err = c.service.AddTodo(ctx, uname, todo)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.DataInvalid):
		util.HttpErrBadRequest(w, err.Error()+": make sure mandatory fields are not empty!")
		return
	case errors.Is(err, errorx.NoDataSaved):
		util.HttpErrBadRequest(w, err.Error()+": there is no data inserted!")
		return
	default:
		util.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	util.HttpCreated(w, "ok", todo)
}

func (c *controller) putTodoItem(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}

func (c *controller) deleteTodoItem(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}
