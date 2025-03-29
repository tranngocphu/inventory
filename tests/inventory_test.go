package tests

import (
	"bytes"
	"encoding/json"
	"inventory/internal/inventory"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateItem(t *testing.T) {
	// Mock repo (simplified for brevity)
	repo := &inventory.Repository{} // Add mock DB later
	handler := inventory.NewHandler(repo)

	item := inventory.Item{Name: "TestItem", Stock: 10}
	body, _ := json.Marshal(item)
	req, _ := http.NewRequest("POST", "/items", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	handler.CreateItem(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)
}
