package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/config/getenv"
	"github.com/felipeazsantos/deploy-gcloudrun-fullcycle-lab01/internal/api"
)

func main() {
	if ok := getenv.LoadConfig("../../.env", ".env"); !ok {
		log.Fatal("failure to load env")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.FindTemperatureByCEP)

	fmt.Println("starting server on port: 8080")
	http.ListenAndServe(":8080", mux)
}
