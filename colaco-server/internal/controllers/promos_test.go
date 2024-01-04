package controllers

import (
	"colaco-server/internal/db"
	"context"
	"net/http/httptest"
	"testing"
)

// Assuming that the Promo, PromoCreateRequest structures are defined appropriately
// and db.MockColacoDB properly mocks the required database operations.

func TestPromosGetAll(t *testing.T) {
	// Initialize mock database and controller
	mockDB := &db.MockColacoDB{}
	ctrl := PromosController{}

	// Add test data to the mock database
	mockDB.Promos = map[string]any{
		"1": Promo{ID: "1", SodaID: "1", Price: 2.0},
		// Add more promos as needed
	}

	// Create a test request with context
	req := httptest.NewRequest("GET", "/promos", nil)
	ctx := context.WithValue(req.Context(), "db", mockDB)
	req = req.WithContext(ctx)

	// Test the GetAll method
	promos, err := ctrl.GetAll(req)
	if err != nil {
		t.Errorf("GetAll returned an error: %v", err)
	}
	if len(promos) != len(mockDB.Promos) {
		t.Errorf("Expected %d promos, got %d", len(mockDB.Promos), len(promos))
	}
}

func TestPromosGetOneById(t *testing.T) {
	// Initialize mock database and controller
	mockDB := &db.MockColacoDB{}
	ctrl := PromosController{}

	// Add test data to the mock database
	mockDB.Promos = map[string]any{
		"1": Promo{ID: "1", SodaID: "1", Price: 2.0},
		"2": Promo{ID: "2", SodaID: "2", Price: 1.5},
	}

	// Create a test request with context
	req := httptest.NewRequest("GET", "/promos/1", nil)
	ctx := context.WithValue(req.Context(), "db", mockDB)
	req = req.WithContext(ctx)

	// Test the GetOneById method
	promo, err := ctrl.GetOneById("1", req)
	if err != nil {
		t.Errorf("GetOneById returned an error: %v", err)
	}

	// Assert the expected outcome
	if promo.ID != "1" {
		t.Errorf("Expected promo ID 1, got %v", promo.ID)
	}
	if promo.Price != 2.0 {
		t.Errorf("Expected promo price 2.0, got %v", promo.Price)
	}
}
