package routes

import (
	"colaco-server/internal/controllers"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
		if !ok {
			zap.L().Error("could not get controller")
			render.Status(r, http.StatusInternalServerError)
			return
		}

		// Get the x-auth-token header value
		uid := r.Header.Get("x-auth-token")
		if uid == "" {
			zap.L().Error("could not get token")
			render.Render(w, r, ErrUnauthorized(errors.New("could not get token")))
		} else {
			isAdmin, err := ctrl.GetIsAdmin(uid, r)
			if err != nil {
				zap.L().Error("could not get user role", zap.String("message", err.Error()))
				render.Render(w, r, ErrUnauthorized(errors.New("could not get user role")))
			}
			if !isAdmin {
				render.Render(w, r, ErrUnauthorized(errors.New("user is not admin")))
			} else {
				next.ServeHTTP(w, r)
			}
		}
	})
}

func authenticatedUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the x-auth-token header value
		uid := r.Header.Get("x-auth-token")
		if uid == "" {
			zap.L().Error("could not get token")
			render.Render(w, r, ErrUnauthorized(errors.New("could not get token")))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func MakeColaCoV1Router(r chi.Router) {
	r.Use(controllers.UsersControllerContext(&controllers.UsersController{}))
	r.Use(controllers.SodasControllerContext(&controllers.SodasController{}))
	r.Use(controllers.PromosControllerContext(&controllers.PromosController{}))

	r.Route("/healthz", func(r chi.Router) {
		r.Get("/", Healthz)
	})

	r.Route("/sodas", func(r chi.Router) {
		ss := &SodasService{}

		r.Get("/", ss.GetAll)
		r.Route("/{sodaID}", func(r chi.Router) {
			r.Get("/", ss.GetOne)
			r.Get("/price", ss.GetOne)
			r.Put("/sell", ss.Sell)

			r.Route("/", func(r chi.Router) {
				r.Use(adminOnly)
				r.Put("/restock", ss.Restock)
			})
		})
	})

	r.Route("/users", func(r chi.Router) {
		us := &UsersService{}

		r.Post("/", us.Create)

		r.Route("/me", func(r chi.Router) {
			r.Use(authenticatedUser)
			r.Get("/", us.Me)
			r.Get("/is_admin", us.IsAdmin)
			r.Get("/balance", us.Balance)
			r.Put("/deposit", us.Deposit)
			r.Put("/withdraw", us.Withdraw)
		})
	})

	r.Route("/promos", func(r chi.Router) {
		ps := &PromosService{}

		// Admins only!
		r.Use(adminOnly)

		r.Get("/", ps.GetAll)
		r.Post("/", ps.Create)
		r.Route("/{promoID}", func(r chi.Router) {
			r.Get("/", ps.GetOne)
		})
	})
}
