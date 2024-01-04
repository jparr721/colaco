package controllers

import (
	"errors"
	"net/http"
)

type Promo struct {
	ID        string  `db:"id"`
	StartDate string  `db:"start_date"`
	EndDate   string  `db:"end_date"`
	SodaID    string  `db:"soda_id"`
	CreatedAt string  `db:"created_at"`
	UpdatedAt string  `db:"updated_at"`
	Price     float64 `db:"price"`
}

type PromoCreateRequest struct {
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	SodaID    string  `json:"soda_id"`
	Price     float64 `json:"price"`
}

// Bind implements render.Binder.
func (p *PromoCreateRequest) Bind(r *http.Request) error {
	// Check if start date is valid
	if p.StartDate == "" {
		return errors.New("start date is required")
	}

	// Check if end date is valid
	if p.EndDate == "" {
		return errors.New("end date is required")
	}

	// Check if price is valid
	if p.Price < 0 {
		return errors.New("price must be greater than or equal to 0")
	}

	if p.SodaID == "" {
		return errors.New("soda id is required")
	}

	return nil
}
