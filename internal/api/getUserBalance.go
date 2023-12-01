package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

type userPayload struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func GetUserBalance(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var payload userPayload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("could not decode JSON: %s\n", err)
		return
	}

	userBalance, err := storage.UserBalance(payload.Name, payload.Surname)
	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, string(userBalance))
	fmt.Println("Endpoint hit: get-user-balance")
}
