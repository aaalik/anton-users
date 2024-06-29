package routes

import (
	"github.com/aaalik/anton-users/internal/middleware"
	"github.com/go-chi/chi"
)

func (rh *RouteHandler) user() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.RequireAuth)
		r.Post("/", rh.UserHandler.CreateUser)
		r.Put("/{userID}", rh.UserHandler.UpdateUser)
		r.Delete("/", rh.UserHandler.DeleteUser)
		r.Post("/list", rh.UserHandler.ListUser)
		r.Get("/{userID}", rh.UserHandler.DetailUser)
	})

	return r
}
