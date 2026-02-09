package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (c *controller) todoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getTodo(w, r)
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

func (c *controller) getTodo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		c.getTodoList(w, r)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid 'id' value")
		return
	}

	data, err := c.service.GetTodoById(r.Context(), int32(idInt))
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get todo")
		return
	}

	if data == nil {
		util.HttpErrNotFound(w, fmt.Sprintf("no data with 'id = %d'", idInt))
		return
	}

	util.HttpOk(w, "ok", data)
}

func (c *controller) getTodoList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	uname, err := claims.GetIssuer()
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get claims context")
		return
	}

	list, err := c.service.GetTodoByUsername(ctx, uname)
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get todo list")
		return
	}

	if len(list) == 0 {
		util.HttpErrNotFound(w, "no data")
		return
	}

	util.HttpOk(w, "ok", list)
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
