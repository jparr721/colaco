package controllers

import (
	"colaco-server/internal/db"
	"context"
	"errors"
	"net/http"
)

type SodasController struct{}

func SodasControllerContext(ctrl *SodasController) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "SodasController", ctrl)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (s *SodasController) GetAll(r *http.Request) ([]Soda, error) {
	db, ok := r.Context().Value("db").(*db.ColacoDB)
	if !ok {
		return nil, errors.New("could not get database connection")
	}

	var sodas []Soda
	err := db.Get("SELECT * FROM sodas", &sodas)
	if err != nil {
		return nil, err
	}

	return sodas, nil
}

func (s *SodasController) GetOneById(id string, r *http.Request) (Soda, error) {
	var soda Soda
	db, ok := r.Context().Value("db").(*db.ColacoDB)
	if !ok {
		return soda, errors.New("could not get database connection")
	}

	err := db.GetOne("SELECT * FROM sodas WHERE id = $1", &soda, id)
	if err != nil {
		return soda, errors.New("could not get soda")
	}

	return soda, nil
}

func (s *SodasController) ChangeStockById(id string, amount int, r *http.Request) (Soda, error) {
	var soda Soda
	db, ok := r.Context().Value("db").(*db.ColacoDB)
	if !ok {
		return soda, errors.New("could not get database connection")
	}

	soda, err := s.GetOneById(id, r)
	if err != nil {
		return soda, err
	}

	// Can we make this change? If not, reject.
	if soda.CurrentQuantity+amount < 0 || soda.CurrentQuantity+amount > soda.MaxQuantity {
		return soda, errors.New("could not change stock")
	}

	// Otherwise, make the update
	err = db.Update("UPDATE sodas SET current_quantity = current_quantity + $1 WHERE id = $2", amount, id)
	if err != nil {
		return soda, errors.New("could not update soda")
	}

	// Update the existing record and return
	soda.CurrentQuantity = soda.CurrentQuantity + amount
	return soda, nil
}
