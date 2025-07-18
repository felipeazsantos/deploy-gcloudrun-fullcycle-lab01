package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env", ".env"); err != nil {
		log.Fatal("failure to load env:", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.FindTemperatureByCEP)

	fmt.Printf("starting server on port: 8080")
	http.ListenAndServe(":8080", mux)
}
