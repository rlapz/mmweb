package controller

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

func (c *controllerAuth) loginHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	uname, pswd, ok := r.BasicAuth()
	if !ok {
		util.HttpErrBadRequest(w, "invalid username or password")
		return
	}

	err := c.service.AuthUser(r.Context(), uname, pswd)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.AuthInvalidCredential):
		util.HttpErrUnauthorized(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to authenticate credential")
		return
	}

	token, err := util.JwtMakeSignedToken(c.signMethod, c.signKey, uname, c.loginExp)
	if err != nil {
		util.HttpErrInternal(w, err, "failed to make and sign token")
		return
	}

	util.HttpOk(w, "ok", token)
}

func (c *controllerAuth) logoutHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	ctx := r.Context()
	claims := util.ContextGetJwtClaims(ctx)
	err := c.service.AuthTokenAdd(ctx, claims["token"].(string))
	if err != nil {
		util.HttpErrInternal(w, err, "failed to invalidate token")
		return
	}

	util.HttpCreated(w, "ok", nil)
}
