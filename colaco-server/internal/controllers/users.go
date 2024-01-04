package controllers

import (
	"colaco-server/internal/db"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type UsersController struct{}

func UsersControllerContext(ctrl *UsersController) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "UsersController", ctrl)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (u *UsersController) GetBalance(id string, r *http.Request) (float64, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return 0, errors.New("could not get database connection")
	}

	var balance UserBalanceResponse
	err := db.GetOne("SELECT balance FROM users WHERE id = $1", &balance, id)
	if err != nil {
		return 0, err
	}

	return balance.Balance, nil
}

func (u *UsersController) GetIsAdmin(id string, r *http.Request) (bool, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return false, errors.New("could not get database connection")
	}

	var isAdmin UserIsAdminResponse
	err := db.GetOne("SELECT is_admin FROM users WHERE id = $1", &isAdmin, id)
	if err != nil {
		return false, err
	}

	return isAdmin.IsAdmin, nil
}

func (u *UsersController) CreateUser(r *http.Request) (User, error) {
	var user User
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return user, errors.New("could not get database connection")
	}

	id, err := db.Create("INSERT INTO users (balance, is_admin) VALUES ($1, $2)", 10, true)
	if err != nil {
		return user, errors.New("could not create user")
	}

	fmt.Println("CREATED ID", id)

	err = db.GetOne("SELECT * FROM users WHERE id = $1", &user, id)
	if err != nil {
		return user, nil
	}

	return user, nil
}

func (u *UsersController) ChangeBalance(amount float64, id string, r *http.Request) (float64, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return -1, errors.New("could not get database connection")
	}

	var balanceResp UserBalanceResponse
	err := db.GetOne("SELECT balance FROM users WHERE id = $1", &balanceResp, id)
	if err != nil {
		return -1, errors.New("could not get user")
	}

	balance := balanceResp.Balance

	if balance+amount < 0 {
		return balance, errors.New("insufficient funds")
	}

	balance += amount

	err = db.Update("UPDATE users SET balance = $1 WHERE id = $2", balance, id)
	if err != nil {
		return -1, errors.New("could not update user")
	}

	return balance, nil
}

func (u *UsersController) GetMe(id string, r *http.Request) (User, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return User{}, errors.New("could not get database connection")
	}

	var user User
	err := db.GetOne("SELECT * FROM users WHERE id = $1", &user, id)
	if err != nil {
		return user, errors.New("could not get user")
	}

	return user, nil
}
