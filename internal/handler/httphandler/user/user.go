package user

import (
	"net/http"

	"github.com/aaalik/anton-users/internal/middleware"
	"github.com/aaalik/anton-users/internal/service"
	"github.com/aaalik/anton-users/internal/utils/validator"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  *service.RequestCreateUser
		resp *service.ResponseUser
		ctx  = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		uh.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	if err := validator.ValidateByName(req, "create_user"); err != nil {
		uh.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := uh.uu.CreateUser(ctx, req)
	if err != nil {
		uh.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	resp = &service.ResponseUser{User: user}
	uh.httpRes.Yay(w, r, http.StatusCreated, resp)
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  *service.RequestUpdateUser
		resp *service.ResponseUser
		ctx  = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		uh.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	req.Id = chi.URLParam(r, "userID")

	if err := validator.ValidateByName(req, "update_user"); err != nil {
		uh.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	user, err := uh.uu.UpdateUser(ctx, req)
	if err != nil {
		uh.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	resp = &service.ResponseUser{User: user}
	uh.httpRes.Yay(w, r, http.StatusCreated, resp)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	err := uh.uu.DeleteUser(ctx, chi.URLParam(r, "userID"))
	if err != nil {
		uh.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	uh.httpRes.Yay(w, r, http.StatusNoContent, nil)
}

func (uh *UserHandler) DetailUser(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *service.ResponseUser
		userID = chi.URLParam(r, "userID")
		ctx    = r.Context()
	)

	if userID == "" {
		userID = middleware.GetUserIdFromToken(r)
	}

	user, err := uh.uu.DetailUser(ctx, userID)
	if err != nil {
		uh.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	resp = &service.ResponseUser{User: user}
	uh.httpRes.Yay(w, r, http.StatusOK, resp)
}

func (uh *UserHandler) ListUser(w http.ResponseWriter, r *http.Request) {
	var (
		req  *service.RequestListUser
		resp *service.ResponseListUser
		ctx  = r.Context()
	)

	if err := render.Decode(r, &req); err != nil {
		uh.httpRes.Nay(w, r, http.StatusBadRequest, err)
		return
	}

	users, count, err := uh.uu.ListUser(ctx, req)
	if err != nil {
		uh.httpRes.Nay(w, r, http.StatusInternalServerError, err)
		return
	}

	resp = &service.ResponseListUser{Users: users, Stats: &service.ResponseStats{Total: count}}
	uh.httpRes.Yay(w, r, http.StatusOK, resp)
}
