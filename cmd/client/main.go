package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "8080"
	}
	res, err := http.DefaultClient.Get("http://" + host + ":" + port + "/number")
	if err != nil {
		log.Panicf("Erro ao realizar a request: %v", err)
	}
	defer res.Body.Close()
	body := make(map[string]int, 1)
	err = json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		log.Panicf("Erro ao decodificar resposta: %v", err)
	}
	log.Printf("Resposta do servidor: %v", body["number"])
}
