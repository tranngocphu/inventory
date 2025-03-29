package tests

import (
	"bytes"
	"encoding/json"
	"inventory/internal/inventory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	// Setup mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	// Mock repo (simplified for brevity)
	repo := inventory.NewRepository(sqlxDB)
	handler := inventory.NewHandler(repo)

	// Expect the INSERT query
	mock.ExpectExec("INSERT INTO items \\(name, stock\\) VALUES \\(\\$1, \\$2\\)").
		WithArgs("TestItem", 10).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Simulate successful insert

	// Prepare request
	item := inventory.Item{Name: "TestItem", Stock: 10}
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	// Call the handler
	handler.CreateItem(rr, req)

	// Assertions
	assert.Equal(t, http.StatusCreated, rr.Code)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Missed expectations: %v", err)
	}
}
