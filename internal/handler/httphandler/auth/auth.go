package auth

import (
	"net/http"

	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/pkg/validator"
	"github.com/go-chi/render"
)

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		req *service.RequestLogin
		ctx = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateByName(req, "login"); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	resp, err := ah.au.Login(ctx, req.Username, req.Password)
	if err != nil {
		ah.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	ah.httpRes.Yay(w, r, http.StatusOK, resp)
}

func (ah *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var (
		req  *service.RequestRefreshToken
		resp *service.ResponseRefreshToken
		ctx  = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateByName(req, "refresh_token"); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	token, exp, err := ah.au.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		ah.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	resp = &service.ResponseRefreshToken{
		Token:     token,
		ExpiresAt: exp,
	}

	ah.httpRes.Yay(w, r, http.StatusOK, resp)
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var (
		req *service.RequestRegister
		ctx = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateByName(req, "register"); err != nil {
		ah.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	resp, err := ah.au.Register(ctx, req)
	if err != nil {
		ah.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	ah.httpRes.Yay(w, r, http.StatusCreated, resp)
}
