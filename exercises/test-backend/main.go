package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type saludoResponse struct {
	Mensaje string `json:"mensaje"`
}

func holaHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(saludoResponse{Mensaje: "Hola AZAC"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hola", holaHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API escuchando en :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
