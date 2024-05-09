package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	res, err := http.DefaultClient.Get("http://localhost:8080/number")
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
