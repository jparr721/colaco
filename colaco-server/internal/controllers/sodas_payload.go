package controllers

import (
	"errors"
	"net/http"
)

type Soda struct {
	ID              string  `json:"id"`
	ProductName     string  `json:"product_name"`
	Description     string  `json:"description,omitempty"`
	Cost            float64 `json:"cost"`
	CurrentQuantity int     `json:"current_quantity"`
	MaxQuantity     int     `json:"max_quantity"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
}

type SodaUpdateRequest struct {
	ID              string   `json:"id"`
	ProductName     *string  `json:"product_name,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Cost            *float64 `json:"cost,omitempty"`
	CurrentQuantity *int     `json:"current_quantity,omitempty"`
	MaxQuantity     *int     `json:"max_quantity,omitempty"`
	CreatedAt       *string  `json:"created_at,omitempty"`
	UpdatedAt       *string  `json:"updated_at,omitempty"`
}

type SodaStockChangeRequest struct {
	Amount int `json:"amount"`
}

// Bind implements render.Bind, to post-process check the requisite fields after unmarshaling
func (s *SodaStockChangeRequest) Bind(r *http.Request) error {
	// Make sure all fields are present, returning an error otherwise
	if s.Amount == 0 || s.Amount < 0 {
		return errors.New("missing field `amount` or insufficient quantity supplied")
	}

	return nil
}
