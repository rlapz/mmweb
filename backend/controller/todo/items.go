package todo

import (
	"net/http"
	"strconv"

	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/model/api"
)

func (t *Todo) itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		t.getItemList(w, r)
	default:
		api.HttpMethodCheck(w, r, "invalid")
	}
}

func (t *Todo) getItemList(w http.ResponseWriter, r *http.Request) {
	query := api.RequestQueryParse(r)

	id := query.Id
	if id == "" {
		api.HttpErrBadRequest(w, "no 'id' parameter found")
		return
	}

	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		api.HttpErrBadRequest(w, "invalid 'id'")
		return
	}

	list, err := t.service.TodoGetItemList(r.Context(), int32(idInt))
	if err != nil {
		api.HttpErrInternal(w, err, "failed to get todo items")
		return
	}

	if len(list) == 0 {
		api.HttpErrNotFound(w, "no data found")
		return
	}

	/* TODO */
	pag := model.Pagination{}
	api.HttpOkWithPagination(w, "ok", list, &pag)
}
