package controller

import (
	"errors"
	"net/http"

	"github.com/rlapz/mmweb/errorx"
	"github.com/rlapz/mmweb/util"
)

func (c *Controller) loginHandler(w http.ResponseWriter, r *http.Request) {
	if !util.HttpMethodCheck(w, r, http.MethodPost) {
		return
	}

	uname, pswd, ok := r.BasicAuth()
	if !ok {
		util.HttpErrBadRequest(w, "invalid username or password")
		return
	}

	err := c.service.Auth(r.Context(), uname, pswd)
	switch {
	case err == nil: // ok
	case errors.Is(err, errorx.AuthInvalidCredential):
		util.HttpErrUnauthorized(w, err.Error())
		return
	default:
		util.HttpErrInternal(w, err, "failed to authenticate credential")
		return
	}

	claims := util.JwtMakeClaims(c.signMethod, uname, c.loginExp)

	signed, err := claims.SignedString(c.signKey)
	if err != nil {
		util.HttpErrInternal(w, err, "failed to sign key")
		return
	}

	util.HttpOk(w, "ok", signed)
}
