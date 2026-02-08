package controller

import (
	"log"
	"net/http"

	"github.com/rlapz/mmweb/util"
)

func (c *controller) indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("index: %p\n", c)
	if r.URL.Path != "/" {
		util.HttpErrNotFound(w, "")
		return
	}

	util.HttpOk(w, "ok", map[string]int{"hello": 1})
}
