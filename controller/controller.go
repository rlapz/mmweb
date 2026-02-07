package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rlapz/mmweb/config"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/service"
)

type controllerItem struct {
	path    string
	handler http.HandlerFunc
}

type Controller struct {
	middleware *middleware.Middleware
	service    *service.Service
	signMethod *jwt.SigningMethodHMAC
	signKey    []byte
	loginExp   time.Duration
}

func Init(cfg *config.Config, mid *middleware.Middleware, srv *service.Service) {
	c := new(Controller)

	c.middleware = mid
	c.service = srv
	c.signMethod = cfg.JwtSignMethod
	c.signKey = cfg.JwtSignatureKey
	c.loginExp = cfg.LoginExp

	contItems := c.handlerList()
	for i := range contItems {
		itm := &contItems[i]
		c.middleware.HandleFunc(itm.path, itm.handler)

		log.Println("route:", itm.path)
	}
}

func (c *Controller) handlerList() []controllerItem {
	return []controllerItem{
		{path: "/", handler: c.indexHandler},
		{path: "/login", handler: c.loginHandler},
		{path: "/todo", handler: c.todoHandler},
	}
}
