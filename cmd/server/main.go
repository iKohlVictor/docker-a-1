package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
)

func main() {
	mux := chi.NewRouter()
	mux.Get("/number", func(w http.ResponseWriter, r *http.Request) {
		randomInt := rand.Int()
		w.Header().Set("Content-Type", "application/json")
		log.Printf("Generated live demo random number: %d", randomInt)
		json.NewEncoder(w).Encode(map[string]int{"number": randomInt})
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Printf("Server listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("Server error: %v", err)
		}
		log.Println("Server stopped accepting new request ...")
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan

	log.Println("Shutting down server gracefully...")
	server.SetKeepAlivesEnabled(false)

	shutdownCtx, shutdownCtxCancel := context.WithCancel(context.Background())
	defer shutdownCtxCancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Panicf("Server shutdown error: %v", err)
	}

	log.Println("Server shutdown successfully!")
}
