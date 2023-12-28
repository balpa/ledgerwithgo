package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github/balpa/ledgerwithgo/internal/storage"
)

type balanceLogUserPayload struct {
	Name      string    `json:"Name"`
	Surname   string    `json:"Surname"`
	StartDate time.Time `json:"StartDate"`
	EndDate   time.Time `json:"EndDate"`
}

func GetBalanceLogByDate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var payload balanceLogUserPayload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("could not decode JSON: %s\n", err)
		return
	}

	storage.GetBalanceLog(payload.Name, payload.Surname, payload.StartDate, payload.EndDate)

	fmt.Println("Endpoint hit: get balance log")
}
