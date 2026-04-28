package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/model/api"
	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"
)

type Todo struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, serv *service.Service) {
	t := new(Todo)
	t.service = serv

	mid.AddHandler("/todo", t.handler, middleware.FLAG_AUTH)
	mid.AddHandler("/todo/item", t.itemHandler, middleware.FLAG_AUTH)
	mid.AddHandler("/todo/items", t.itemsHandler, middleware.FLAG_AUTH)
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
		api.HttpMethodCheck(w, r, "invalid")
	}
}

func (t *Todo) get(w http.ResponseWriter, r *http.Request) {
	query := api.RequestQueryParse(r)

	id := query.Id
	if id == "" {
		t.getList(w, r)
		return
	}

	idNum, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		api.HttpErrBadRequest(w, "invalid id")
		return
	}

	todo, err := t.service.TodoGet(r.Context(), int32(idNum))
	switch {
	case err == nil:
	case errors.Is(err, errorx.NoDataFound):
		api.HttpErrNotFound(w, err.Error())
		return
	default:
		api.HttpErrInternal(w, err, "failed to get 'todo'")
		return
	}

	api.HttpOk(w, "ok", todo)
}

func (t *Todo) getList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := api.RequestQueryParse(r)
	userId := t.service.AuthContextGetUserId(ctx)
	pag := query.Paginate()
	list, err := t.service.TodoGetList(ctx, userId, pag)
	if err != nil {
		api.HttpErrInternal(w, err, "failed to get todo list")
		return
	}

	api.HttpOkWithPagination(w, "ok", list, api.ResponsePaginationNew(pag))
}

func (t *Todo) post(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, err := util.ParseJsonReader[model.Todo](r.Body)
	if err != nil {
		api.HttpErrBadRequest(w, "invalid body")
		return
	}

	todo.IdUser = t.service.AuthContextGetUserId(ctx)
	err = t.service.TodoAdd(ctx, todo)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.DataExists), errors.Is(err, errorx.DataInvalid):
		api.HttpErrBadRequest(w, err.Error())
		return
	default:
		api.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	api.HttpCreated(w, "ok", nil)
}

func (t *Todo) put(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	todo, err := util.ParseJsonReader[model.Todo](r.Body)
	if err != nil {
		api.HttpErrBadRequest(w, "invalid body")
		return
	}

	todo.IdUser = t.service.AuthContextGetUserId(ctx)
	err = t.service.TodoUpdate(ctx, todo)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.NoDataUpdated), errors.Is(err, errorx.DataExists):
		api.HttpErrBadRequest(w, err.Error())
		return
	default:
		api.HttpErrInternal(w, err, "failed to update todo")
		return
	}

	api.HttpCreated(w, "ok", nil)
}

func (t *Todo) delete(w http.ResponseWriter, r *http.Request) {
	query := api.RequestQueryParse(r)

	id := query.Id
	if id == "" {
		t.getList(w, r)
		return
	}

	idNum, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		api.HttpErrBadRequest(w, "invalid id")
		return
	}

	err = t.service.TodoDelete(r.Context(), int32(idNum))
	switch {
	case err == nil:
	case errors.Is(err, errorx.NoDataFound):
		api.HttpErrNotFound(w, err.Error())
		return
	default:
		api.HttpErrInternal(w, err, "failed to delete 'todo'")
		return
	}

	api.HttpOk(w, "ok", nil)
}
