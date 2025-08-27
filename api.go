package main

import (
	"encoding/json"
	"net/http"
)

type Pessoa struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

func main() {
	http.HandleFunc("/pessoas", func(w http.ResponseWriter, r *http.Request) {
		// lista de pessoas simulada
		pessoas := []Pessoa{
			{Nome: "Ana", Idade: 25},
			{Nome: "Bruno", Idade: 30},
			{Nome: "Carla", Idade: 22},
			{Nome: "Diego", Idade: 28},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(pessoas)
	})

	http.ListenAndServe(":8080", nil)
}
