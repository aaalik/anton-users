package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	chim "github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"

	apimiddleware "github.com/aaalik/anton-users/internal/middleware"
)

func New(log *logrus.Logger, routeHandler *RouteHandler) *Route {
	routeHandler.log = log
	return &Route{
		log:          log,
		routeHandler: routeHandler,
	}
}

func (ir Route) Init() *chi.Mux {
	r := chi.NewRouter()
	r.Use(
		chim.NoCache,
		chim.RedirectSlashes,
		chim.Heartbeat("/ping"),
		chim.RequestID,
		chim.Recoverer,
		chim.RealIP,
		apimiddleware.RequestLogger(ir.log),
		apimiddleware.CORS,
	)
	r.Route("/v1", func(r chi.Router) {
		r.Mount("/user", ir.routeHandler.user())
		r.Post("/login", ir.routeHandler.AuthHandler.Login)
		r.Post("/refresh-token", ir.routeHandler.AuthHandler.RefreshToken)
		r.Post("/register", ir.routeHandler.AuthHandler.Register)
	})
	r.Group(func(r chi.Router) {
		r.Use(apimiddleware.RequireAuth)
		r.Get("/me", ir.routeHandler.UserHandler.DetailUser)
	})

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		ir.log.Infof("%s %s", method, route)
		return nil
	}

	if err := chi.Walk(r, walkFunc); err != nil {
		ir.log.Panicln(err)
	}

	return r
}
