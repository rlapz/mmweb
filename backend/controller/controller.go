package controller

import (
	"github.com/rlapz/mmweb/controller/auth"
	"github.com/rlapz/mmweb/controller/todo"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/service"
)

type controller struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, srv *service.Service) {
	c := new(controller)
	c.service = srv

	mid.AddHandler("/", c.indexHandler, 0)

	auth.Init(mid, srv)
	todo.Init(mid, srv)
}
