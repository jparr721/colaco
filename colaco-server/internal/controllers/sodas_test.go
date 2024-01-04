package controllers

import (
	"colaco-server/internal/db"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Helper function to create a test request and context.
func createTestRequest() *http.Request {
	req := httptest.NewRequest("GET", "/", nil)
	ctx := context.Background()
	return req.WithContext(ctx)
}

// TestGetPromoPrice tests the GetPromoPrice method of SodasController.
func TestGetPromoPrice(t *testing.T) {
	// Initialize mock database and controller
	mockDB := &db.MockColacoDB{}
	ctrl := SodasController{}

	// Add test data to the mock database
	// Assuming dbmock.Promo and dbmock.Soda are correctly defined
	mockDB.Sodas = map[string]any{
		"1": Soda{ID: "1", ProductName: "Cola", Cost: 2.5},
		"2": Soda{ID: "2", ProductName: "Lemon Lime", Cost: 2.0},
	}

	mockDB.Promos = map[string]any{
		"1": Promo{ID: "1", SodaID: "1", Price: 2.0}, // Promo for Cola
	}

	// Create a test request
	req := createTestRequest()
	ctx := context.WithValue(req.Context(), "db", mockDB)
	req = req.WithContext(ctx)

	// Test the method for a soda with a promo
	price, err := ctrl.GetPromoPrice("1", req)
	if err != nil {
		t.Errorf("GetPromoPrice returned an error for soda with promo: %v", err)
	}
	if price != 2.0 {
		t.Errorf("GetPromoPrice returned incorrect price for soda with promo, got: %v, want: %v", price, 2.0)
	}

	// Test the method for a soda without a promo
	price, err = ctrl.GetPromoPrice("2", req)
	if err == nil {
		t.Errorf("GetPromoPrice should return an error for soda without promo")
	}
	if price != 0 {
		t.Errorf("GetPromoPrice should return 0 for soda without promo, got: %v", price)
	}
}

// TestGetAll tests the GetAll method of SodasController.
func TestGetAll(t *testing.T) {
	// Initialize mock database and controller
	mockDB := &db.MockColacoDB{}
	ctrl := SodasController{}

	// Add test data to the mock database
	// Assuming dbmock.Promo and dbmock.Soda are correctly defined
	mockDB.Sodas = map[string]any{
		"1": Soda{ID: "1", ProductName: "Cola", Cost: 2.5},
		"2": Soda{ID: "2", ProductName: "Lemon Lime", Cost: 2.0},
	}

	mockDB.Promos = map[string]any{
		"1": Promo{ID: "1", SodaID: "1", Price: 2.0}, // Promo for Cola
	}

	// Create a test request
	req := createTestRequest()
	ctx := context.WithValue(req.Context(), "db", mockDB)
	req = req.WithContext(ctx)

	// Test the method for a soda with a promo
	sodas, err := ctrl.GetAll(req)
	if err != nil {
		t.Errorf("GetPromoPrice returned an error for soda with promo: %v", err)
	}

	if len(sodas) != 2 {
		t.Errorf("GetAll returned incorrect number of sodas %v want: %v", len(sodas), 2)
	}
}

// TestGetOneById tests the GetOneById method of SodasController.
func TestGetOneById(t *testing.T) {
	// Initialize mock database and controller
	mockDB := &db.MockColacoDB{}
	ctrl := SodasController{}

	// Add test data to the mock database
	// Assuming dbmock.Promo and dbmock.Soda are correctly defined
	mockDB.Sodas = map[string]any{
		"1": Soda{ID: "1", ProductName: "Cola", Cost: 2.5},
		"2": Soda{ID: "2", ProductName: "Lemon Lime", Cost: 2.0},
	}

	mockDB.Promos = map[string]any{
		"1": Promo{ID: "1", SodaID: "1", Price: 2.0}, // Promo for Cola
	}

	// Create a test request
	req := createTestRequest()
	ctx := context.WithValue(req.Context(), "db", mockDB)
	req = req.WithContext(ctx)

	// Test the method for a soda with a promo
	soda, err := ctrl.GetOneById("1", req)
	if err != nil {
		t.Errorf("GetPromoPrice returned an error for soda with promo: %v", err)
	}

	if soda.ID != "1" {
		t.Errorf("GetOneById returned incorrect soda %v want: %v", soda.ID, "1")
	}
}
