package routes

import (
	"colaco-server/internal/controllers"

	"github.com/go-chi/chi/v5"
)

func MakeColaCoRouter(r chi.Router) {
	r.Route("/healthz", func(r chi.Router) {
		r.Get("/", Healthz)
	})

	r.Route("/sodas", func(r chi.Router) {
		ss := &SodasService{}
		r.Use(controllers.SodasControllerContext(&controllers.SodasController{}))

		r.Get("/", ss.GetAll)
		r.Route("/{sodaID}", func(r chi.Router) {
			r.Get("/", ss.GetOne)
			r.Get("/price", ss.GetOne)
			r.Put("/restock", ss.Restock)
			r.Put("/sell", ss.Sell)
		})
	})

	r.Route("/users", func(r chi.Router) {
		us := &UsersService{}
		r.Use(controllers.UsersControllerContext(&controllers.UsersController{}))

		r.Post("/", us.Create)

		r.Route("/{userID}", func(r chi.Router) {
			r.Get("/is_admin", us.IsAdmin)
			r.Get("/balance", us.Balance)
			r.Put("/deposit", us.Deposit)
			r.Put("/withdraw", us.Withdraw)
		})
	})
}
