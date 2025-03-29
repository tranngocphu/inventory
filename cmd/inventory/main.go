package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	log.Println(("Starting inventory service on port :8080"))
	log.Fatal(http.ListenAndServe(":8080", r))
}
