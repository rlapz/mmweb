package auth

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/middleware"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/service"
	"github.com/rlapz/mmweb/util"
)

type Auth struct {
	service *service.Service
}

func Init(mid *middleware.Middleware, serv *service.Service) {
	a := new(Auth)
	a.service = serv

	mid.AddHandler("/login", a.login, middleware.FLAG_AUTH_EXCLUDED)
	mid.AddHandler("/logout", a.logout, 0)
	mid.AddHandler("/register", a.register, middleware.FLAG_AUTH_EXCLUDED)
}

func (a *Auth) login(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	uname, pswd, ok := r.BasicAuth()
	if !ok || (uname == "") || (pswd == "") {
		util.HttpErrBadRequest(w, "invalid username or password")
		return
	}

	token, err := a.service.AuthLogin(r.Context(), uname, pswd)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.AuthInvalidCredential):
		util.HttpErrUnauthorized(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to authenticate credential")
		return
	}

	util.HttpOk(w, "ok", token)
}

func (a *Auth) logout(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	err := a.service.AuthLogout(ctx, claims["token"].(string))
	if err != nil {
		util.HttpErrInternal(w, err, "failed to logged out")
		return
	}

	util.HttpCreated(w, "ok", nil)
}

func (a *Auth) register(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	user, err := util.HttpJsonParseBody[model.User](r.Body)
	if err != nil {
		util.HttpErrBadRequest(w, err.Error())
		return
	}

	err = a.service.AuthRegister(ctx, user)
	switch {
	case err == nil:
	case errors.Is(err, errorx.DataInvalid):
		util.HttpErrBadRequest(w, err.Error())
		return
	case errors.Is(err, errorx.DataExists):
		util.HttpErrBadRequest(w, "such user already exists!")
		return
	default:
		util.HttpErrInternal(w, err, "failed to add new item")
		return
	}

	user.Password = "***"
	util.HttpCreated(w, "ok", user)
}
