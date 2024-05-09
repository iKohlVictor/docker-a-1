package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	mux.Get("/number", func(w http.ResponseWriter, r *http.Request) {
		randomInt := rand.Int()
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Generated random number: %d", randomInt)
		json.NewEncoder(w).Encode(map[string]int{"number": randomInt})
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Server listening on %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Panicf("Server error: %v", err)
	}
}
