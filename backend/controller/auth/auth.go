package auth

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/model/api"
	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"
)

type Auth struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, serv *service.Service) {
	a := new(Auth)
	a.service = serv

	mid.AddHandler("/auth/login", a.login, 0)
	mid.AddHandler("/auth/logout", a.logout, middleware.FLAG_AUTH)
	mid.AddHandler("/auth/register", a.register, 0)
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	if !api.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	uname, pswd, ok := r.BasicAuth()
	if !ok || (uname == "") || (pswd == "") {
		api.HttpErrBadRequest(w, "invalid username or password")
		return
	}

	token, err := a.service.AuthLogin(r.Context(), uname, pswd)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.AuthInvalidCredential):
		api.HttpErrUnauthorized(w, err.Error())
		return
	default:
		api.HttpErrInternal(w, err, "failed to authenticate credential")
		return
	}

	api.HttpOk(w, "ok", token)
}

func (a *Auth) logout(w http.ResponseWriter, r *http.Request) {
	if !api.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	token := a.service.AuthContextGetToken(ctx)
	err := a.service.AuthLogout(ctx, token)
	if err != nil {
		api.HttpErrInternal(w, err, "failed to log out")
		return
	}

	api.HttpCreated(w, "ok", nil)
}

func (a *Auth) register(w http.ResponseWriter, r *http.Request) {
	if !api.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	user, err := util.ParseJsonReader[model.User](r.Body)
	if err != nil {
		api.HttpErrBadRequest(w, err.Error())
		return
	}

	err = a.service.AuthRegister(ctx, user)
	switch {
	case err == nil:
	case errors.Is(err, errorx.DataInvalid):
		api.HttpErrBadRequest(w, err.Error())
		return
	case errors.Is(err, errorx.DataExists):
		api.HttpErrBadRequest(w, "such user already exists!")
		return
	default:
		api.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	user.Password = "***"
	api.HttpCreated(w, "ok", user)
}
