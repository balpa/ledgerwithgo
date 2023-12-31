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
	r.HandleFunc("/api/get-all-balances", api.GetAllBalances)
	r.HandleFunc("/api/get-user-balance", api.GetUserBalance)
	r.HandleFunc("/api/transfer-credit", api.TransferCredit)
	r.HandleFunc("/api/balance-log-by-date", api.GetBalanceLogByDate)

	http.ListenAndServe(":8080", nil)
}
