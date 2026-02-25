package controller

import (
	"net/http"

	"github.com/rlapz/mmweb/util"
)

func (c *controller) indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		util.HttpErrNotFound(w, "")
		return
	}

	util.HttpOk(w, "ok", nil)
}
