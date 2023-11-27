package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

type creditPayload struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Amount  int    `json:"amount"`
}

func AddCredit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var creditPayload creditPayload
	err := decoder.Decode(&creditPayload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("could not decode JSON: %s\n", err)
		return
	}

	storage.AddCredit(creditPayload.Name, creditPayload.Surname, creditPayload.Amount)

	fmt.Println("Endpoint hit: add-credit")
}
