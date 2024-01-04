package routes

import (
	"colaco-server/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type PromosService struct{}

func (ps *PromosService) GetAll(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("PromosController").(*controllers.PromosController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	promos, err := ctrl.GetAll(r)
	if err != nil {
		zap.L().Error("failed to get all promos", zap.String("message", err.Error()))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, promos)
}

func (ps *PromosService) GetOne(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("PromosController").(*controllers.PromosController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	promoId := chi.URLParam(r, "promoID")
	if promoId == "" {
		zap.L().Error("Could not get promo id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	promo, err := ctrl.GetOneById(promoId, r)
	if err != nil {
		zap.L().Error("Could not get promo", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, promo)
}

func (ps *PromosService) Create(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("PromosController").(*controllers.PromosController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	data := &controllers.PromoCreateRequest{}
	if err := render.Bind(r, data); err != nil {
		zap.L().Error("could not bind request", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	promo, err := ctrl.Create(data, r)
	if err != nil {
		zap.L().Error("could not create promo", zap.String("message", err.Error()))
		render.Status(r, http.StatusInternalServerError)
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, promo)
}
