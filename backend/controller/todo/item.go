package todo

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (t *Todo) itemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.getItem(w, r)
	case http.MethodPost:
		t.postItem(w, r)
	case http.MethodPut:
		t.putItem(w, r)
	case http.MethodDelete:
		t.deleteItem(w, r)
	default:
		util.HttpMethodCheck(w, r, "invalid")
	}
}

func (t *Todo) getItem(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	todoId := query.Get("id_todo")
	if todoId == "" {
		util.HttpErrBadRequest(w, "no 'id_todo' parameter found")
		return
	}

	id := query.Get("id")
	if id == "" {
		t.getItemList(w, r, todoId)
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid 'id'")
		return
	}

	item, err := t.service.TodoGetItem(r.Context(), int32(idInt))
	switch {
	case err == nil:
	case errors.Is(err, errorx.NoDataFound):
		util.HttpErrNotFound(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to get 'todo item'")
		return
	}

	util.HttpOk(w, "ok", item)
}

func (t *Todo) getItemList(w http.ResponseWriter, r *http.Request, id string) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid 'id_todo'")
		return
	}

	list, err := t.service.TodoGetItemList(r.Context(), int32(idInt))
	if err != nil {
		util.HttpErrInternal(w, err, "failed to get todo items")
		return
	}

	if len(list) == 0 {
		util.HttpErrNotFound(w, "no data found")
		return
	}

	util.HttpOk(w, "ok", list)
}

func (t *Todo) postItem(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	item, err := util.HttpJsonParseBody[model.TodoItem](r.Body)
	if err != nil {
		util.HttpErrBadRequest(w, "invalid body: "+err.Error())
		return
	}

	err = t.service.TodoAddItem(ctx, item)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.NoDataFound):
		util.HttpErrBadRequest(w, "no such todo")
		return
	case errors.Is(err, errorx.DataExists), errors.Is(err, errorx.DataInvalid):
		util.HttpErrBadRequest(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	util.HttpOk(w, "ok", item)
}

func (t *Todo) putItem(w http.ResponseWriter, r *http.Request) {
	_ = r
	util.HttpOk(w, "TODO", nil)
}

func (t *Todo) deleteItem(w http.ResponseWriter, r *http.Request) {
	_ = r
	util.HttpOk(w, "TODO", nil)
}
