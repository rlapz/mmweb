package controller

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/service"
)

type controller struct {
	service *service.Service
}

type controllerAuth struct {
	service    *service.Service
	signMethod *jwt.SigningMethodHMAC
	signKey    []byte
	loginExp   time.Duration
}

func Init(cfg *config.Config, mid *middleware.Middleware, srv *service.Service) {
	c := new(controller)
	c.service = srv

	mid.AddHandler("/", c.indexHandler, middleware.FLAG_AUTH_EXCLUDED)
	mid.AddHandler("/todo", c.todoHandler, 0)
	mid.AddHandler("/blog", c.blogHandler, 0)
}

func InitAuth(cfg *config.Config, mid *middleware.Middleware, srv *service.Service) {
	ca := new(controllerAuth)
	ca.service = srv
	ca.signMethod = cfg.JwtSignMethod
	ca.signKey = cfg.JwtSignatureKey
	ca.loginExp = cfg.LoginExp

	mid.AddHandler("/login", ca.loginHandler, middleware.FLAG_AUTH_EXCLUDED)
	mid.AddHandler("/logout", ca.logoutHandler, 0)
}
