package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"
)

type Todo struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, serv *service.Service) {
	t := new(Todo)
	t.service = serv

	mid.AddHandler("/todo", t.handler, 0)
	mid.AddHandler("/todo/item", t.itemHandler, 0)
}

func (t *Todo) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.get(w, r)
	case http.MethodPost:
		t.post(w, r)
	case http.MethodPut:
		t.put(w, r)
	case http.MethodDelete:
		t.delete(w, r)
	default:
		util.HttpMethodCheck(w, r, "invalid")
	}
}

func (t *Todo) get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		t.getList(w, r)
		return
	}

	idNum, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid id")
		return
	}

	todo, err := t.service.TodoGet(r.Context(), int32(idNum))
	switch {
	case err == nil:
	case errors.Is(err, errorx.NoDataFound):
		util.HttpErrNotFound(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to get 'todo'")
		return
	}

	util.HttpOk(w, "ok", todo)
}

func (t *Todo) getList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	uname, err := claims.GetIssuer()
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get claims context")
		return
	}

	list, err := t.service.TodoGetList(ctx, uname)
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

func (t *Todo) post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, err := util.HttpJsonParseBody[model.Todo](r.Body)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid body")
		return
	}

	claims := util.ContextGetJwtClaims(ctx)
	uname, err := claims.GetIssuer()
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get context claims")
		return
	}

	err = t.service.TodoAdd(ctx, todo, uname)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.DataExists), errors.Is(err, errorx.DataInvalid):
		util.HttpErrBadRequest(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	util.HttpCreated(w, "ok", todo)
}

func (t *Todo) put(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}

func (t *Todo) delete(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}
