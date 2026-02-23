package controller

import (
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/service"
)

type controller struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, srv *service.Service) {
	c := new(controller)
	c.service = srv

	mid.AddHandler("/", c.indexHandler, middleware.FLAG_AUTH_EXCLUDED)
	mid.AddHandler("/login", c.loginHandler, middleware.FLAG_AUTH_EXCLUDED)
	mid.AddHandler("/logout", c.logoutHandler, 0)
	mid.AddHandler("/register", c.registerHandler, middleware.FLAG_AUTH_EXCLUDED)

	mid.AddHandler("/todo", c.todoHandler, 0)
	mid.AddHandler("/blog", c.blogHandler, 0)
}
