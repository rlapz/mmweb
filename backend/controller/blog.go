package controller

import (
	"net/http"

	"github.com/rlapz/mmweb/util"
)

func (c *controller) blogHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.getBlog(w, r)
	case http.MethodPost:
		c.postBlogItem(w, r)
	case http.MethodPut:
		c.putBlogItem(w, r)
	case http.MethodDelete:
		c.deleteBlogItem(w, r)
	default:
		util.HttpMethodCheck(w, r, "invalid")
	}
}

func (c *controller) getBlog(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		c.getBlogList(w, r)
		return
	}

	util.HttpOk(w, "TODO", nil)
}

func (c *controller) getBlogList(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}

func (c *controller) postBlogItem(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}

func (c *controller) putBlogItem(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}

func (c *controller) deleteBlogItem(w http.ResponseWriter, r *http.Request) {
	util.HttpOk(w, "TODO", nil)

	_ = r
}
