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

func (s *SodasController) GetPromoPrice(id string, r *http.Request) (float64, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return 0, errors.New("could not get database connection")
	}

	var promo Promo
	// Get the most recent promo for this soda
	err := db.GetOne("SELECT * FROM promos WHERE soda_id = $1 AND end_date >= CURRENT_DATE ORDER BY start_date", &promo, id)
	if err != nil {
		return 0, err
	}

	return promo.Price, nil
}

func (s *SodasController) GetAll(r *http.Request) ([]Soda, error) {
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return nil, errors.New("could not get database connection")
	}

	var sodas []Soda
	err := db.Get("SELECT * FROM sodas", &sodas)
	if err != nil {
		return nil, err
	}

	for i := range sodas {
		// Get the promo price for this soda, if there is an entry
		promoPrice, err := s.GetPromoPrice(sodas[i].ID, r)
		if err != nil {
			// If the error is just "sql: no rows in result set", then we can ignore it
			if err.Error() == "sql: no rows in result set" {
				continue
			}

			return nil, err
		}

		sodas[i].Cost = promoPrice
	}

	return sodas, nil
}

func (s *SodasController) GetOneById(id string, r *http.Request) (Soda, error) {
	var soda Soda
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
	if !ok {
		return soda, errors.New("could not get database connection")
	}

	err := db.GetOne("SELECT * FROM sodas WHERE id = $1", &soda, id)
	if err != nil {
		return soda, errors.New("could not get soda")
	}

	// Get the promo price for this soda, if there is an entry
	promoPrice, err := s.GetPromoPrice(soda.ID, r)
	if err != nil {
		// If the error is just "sql: no rows in result set", then we can ignore it
		if err.Error() != "sql: no rows in result set" {
			return soda, err
		}
	}

	soda.Cost = promoPrice

	return soda, nil
}

func (s *SodasController) ChangeStockById(id string, amount int, r *http.Request) (Soda, error) {
	var soda Soda
	db, ok := r.Context().Value("db").(db.ColacoDBInterface)
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
