package routes

import (
	"colaco-server/internal/controllers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type UsersService struct{}

func (u *UsersService) Create(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	user, err := ctrl.CreateUser(r)
	if err != nil {
		zap.L().Error("failed to create user", zap.String("message", err.Error()))
		render.Status(r, http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, user)
}

func (u *UsersService) Balance(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	if userID == "" {
		zap.L().Error("Could not get user id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	balance, err := ctrl.GetBalance(userID, r)
	if err != nil {
		zap.L().Error("Could not get user balance", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &controllers.UserBalanceResponse{Balance: balance})
}

func (u *UsersService) IsAdmin(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
	if !ok {
		zap.L().Error("could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	if userID == "" {
		zap.L().Error("Could not get user id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	isAdmin, err := ctrl.GetIsAdmin(userID, r)
	if err != nil {
		zap.L().Error("Could not get user balance", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &controllers.UserIsAdminResponse{IsAdmin: isAdmin})
}

func (u *UsersService) Deposit(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
	if !ok {
		zap.L().Error("Could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	if userID == "" {
		zap.L().Error("Could not get user id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	data := &controllers.UserBalanceUpdateRequest{}
	if err := render.Bind(r, data); err != nil {
		zap.L().Error("Could not bind request", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	balance, err := ctrl.ChangeBalance(data.Amount, userID, r)
	if err != nil {
		zap.L().Error("Could not get user balance", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &controllers.UserBalanceUpdateResponse{NewBalance: balance})
}

func (u *UsersService) Withdraw(w http.ResponseWriter, r *http.Request) {
	ctrl, ok := r.Context().Value("UsersController").(*controllers.UsersController)
	if !ok {
		zap.L().Error("Could not get controller")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	userID := chi.URLParam(r, "userID")
	if userID == "" {
		zap.L().Error("Could not get user id")
		render.Status(r, http.StatusInternalServerError)
		return
	}

	data := &controllers.UserBalanceUpdateRequest{}
	if err := render.Bind(r, data); err != nil {
		zap.L().Error("Could not bind request", zap.String("message", err.Error()))
		render.Render(w, r, ErrBadRequest(err))
		return
	}

	balance, err := ctrl.ChangeBalance(-data.Amount, userID, r)
	if err != nil {
		zap.L().Error("Could not get user balance", zap.String("message", err.Error()))
		render.Render(w, r, ErrRecordNotFound(err))
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, &controllers.UserBalanceUpdateResponse{NewBalance: balance})
}
