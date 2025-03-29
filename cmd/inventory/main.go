package main

import (
	"inventory/internal/db"
	"inventory/internal/inventory"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	db, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	repo := inventory.NewRepository(db)
	handler := inventory.NewHandler(repo)

	// Initialize DB schema
	db.Exec("CREATE TABLE IF NOT EXISTS items (id SERIAL PRIMARY KEY, name TEXT, stock INT)")

	r := mux.NewRouter()
	// Define routes and route handlers
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.HandleFunc("/items", handler.CreateItem).Methods("POST")

	log.Println(("Starting inventory service at http://localhost:8080"))
	log.Fatal(http.ListenAndServe(":8080", r))
}
