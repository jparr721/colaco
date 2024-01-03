package controllers

import (
	"errors"
	"net/http"
)

type User struct {
	ID        string  `json:"id"`
	Balance   float64 `json:"balance"`
	IsAdmin   bool    `json:"is_admin"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type UserCreateResponse struct {
	ID string `json:"id"`
}

type UserBalanceResponse struct {
	Balance float64 `json:"balance"`
}

type UserIsAdminResponse struct {
	IsAdmin bool `json:"is_admin"`
}

type UserBalanceUpdateResponse struct {
	NewBalance int `json:"new_balance"`
}

type UserBalanceUpdateRequest struct {
	Amount int `json:"amount"`
}

// Bind implements render.Bind.
func (u *UserBalanceUpdateRequest) Bind(r *http.Request) error {
	// Make sure all fields are present, returning an error otherwise
	if u.Amount == 0 || u.Amount < 0 {
		return errors.New("missing field `amount` or insufficient quantity supplied")
	}

	return nil
}
