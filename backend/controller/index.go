package controller

import (
	"net/http"

	"github.com/rlapz/mmweb/util"
)

func (c *controller) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		util.HttpErrNotFound(w, "invalid path")
		return
	}

	util.HttpOk(w, "ok", nil)
}
