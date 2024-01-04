package controllers

import (
	"colaco-server/internal/db"
	"context"
	"errors"
	"fmt"
	"net/http"
)

type PromosController struct{}

func PromosControllerContext(ctrl *PromosController) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), "PromosController", ctrl)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func (p *PromosController) GetAll(r *http.Request) ([]Promo, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return nil, errors.New("could not get database connection")
	}

	var promos []Promo
	err := db.Get("SELECT * FROM promos", &promos)
	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (p *PromosController) GetOneById(id string, r *http.Request) (Promo, error) {
	var promo Promo
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return promo, errors.New("could not get database connection")
	}

	err := db.GetOne("SELECT * FROM promos WHERE id = $1", &promo, id)
	if err != nil {
		return promo, err
	}

	return promo, nil
}

func (p *PromosController) GetAllBySodaId(id string, r *http.Request) ([]Promo, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return nil, errors.New("could not get database connection")
	}

	var promos []Promo
	err := db.Get("SELECT * FROM promos WHERE soda_id = $1", &promos, id)
	if err != nil {
		return nil, err
	}

	return promos, nil
}

func (p *PromosController) Create(data *PromoCreateRequest, r *http.Request) (Promo, error) {
	var promo Promo
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return promo, errors.New("could not get database connection")
	}

	id, err := db.Create("INSERT INTO promos (start_date, end_date, soda_id, price) VALUES ($1, $2, $3, $4)", data.StartDate, data.EndDate, data.SodaID, data.Price)
	if err != nil {
		return promo, err
	}

	fmt.Println("GENERATED ID", id)

	promo, err = p.GetOneById(id, r)
	if err != nil {
		return promo, err
	}

	return promo, nil
}
