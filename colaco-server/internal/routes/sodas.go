package routes

import (
	"colaco-server/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type SodasService struct{}

func (s *SodasService) GetAll(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("SodasController").(*controllers.SodasController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	sodas, err := ctrl.GetAll(r)
	if err != nil {
		zap.L().Error("failed to get all sodas", zap.String("message", err.Error()))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, sodas)
}

func (s *SodasService) GetOne(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("SodasController").(*controllers.SodasController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	sodaId := chi.URLParam(r, "sodaID")
	if sodaId == "" {
		zap.L().Error("Could not get soda id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	soda, err := ctrl.GetOneById(sodaId, r)
	if err != nil {
		zap.L().Error("Could not get soda", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, soda)
}

func (s *SodasService) Restock(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("SodasController").(*controllers.SodasController)
	if !ok {
		zap.L().Error("Could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	sodaId := chi.URLParam(r, "sodaID")
	if sodaId == "" {
		zap.L().Error("Could not get soda id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	data := &controllers.SodaStockChangeRequest{}
	if err := render.Bind(r, data); err != nil {
		zap.L().Error("Could not bind request", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	soda, err := ctrl.ChangeStockById(sodaId, data.Amount, r)
	if err != nil {
		zap.L().Error("Could not get soda", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, soda)
}

func (s *SodasService) Sell(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("SodasController").(*controllers.SodasController)
	if !ok {
		zap.L().Error("Could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	sodaId := chi.URLParam(r, "sodaID")
	if sodaId == "" {
		zap.L().Error("Could not get soda id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	data := &controllers.SodaStockChangeRequest{}
	if err := render.Bind(r, data); err != nil {
		zap.L().Error("Could not bind request", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	// Negate the amount since we're subtracting it
	soda, err := ctrl.ChangeStockById(sodaId, -data.Amount, r)
	if err != nil {
		zap.L().Error("Could not get soda", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, soda)
}
