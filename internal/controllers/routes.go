package controllers

import (
	"net/http"

	"github/balpa/ledgerwithgo/internal/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func HandleRequests() {
	r := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.Handle("/", corsHandler.Handler(r))

	r.HandleFunc("/api/welcome", api.Welcome)
	r.HandleFunc("/api/create-user", api.CreateUser)
	r.HandleFunc("/api/add-credit", api.AddCredit)

	http.ListenAndServe(":8080", nil)
}
