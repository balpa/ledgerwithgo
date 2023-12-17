package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

type transferCreditPayload struct {
	SenderName      string `json:"SenderName"`
	SenderSurname   string `json:"SenderSurname"`
	ReceiverName    string `json:"ReceiverName"`
	ReceiverSurname string `json:"ReceiverSurname"`
	TransferAmount  int    `json:"TransferAmount"`
}

func TransferCredit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var payload transferCreditPayload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("could not decode JSON: %s\n", err)
		return
	}

	storage.TransferCredit(
		payload.SenderName,
		payload.SenderSurname,
		payload.ReceiverName,
		payload.ReceiverSurname,
		payload.TransferAmount)

	fmt.Println("Endpoint hit: transfer credit")
}
