package controller

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/model"
	"github.com/rlapz/mmweb/util"
)

func (c *controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	uname, pswd, ok := r.BasicAuth()
	if !ok {
		util.HttpErrBadRequest(w, "invalid username or password")
		return
	}

	token, err := c.service.AuthLogin(r.Context(), uname, pswd)
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

func (c *controller) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	err := c.service.AuthLogout(ctx, claims["token"].(string))
	if err != nil {
		util.HttpErrInternal(w, err, "failed to logged out")
		return
	}

	util.HttpCreated(w, "ok", nil)
}

func (c *controller) registerHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	user, err := util.HttpJsonParseBody[model.User](r.Body)
	if err != nil {
		util.HttpErrBadRequest(w, "failed to parse request body")
		return
	}

	err = c.service.AuthRegister(r.Context(), user)
	switch {
	case err == nil:
	case errors.Is(err, errorx.DataInvalid):
		util.HttpErrBadRequest(w, "make sure mandatory fields are not empty!")
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
