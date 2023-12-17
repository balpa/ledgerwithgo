package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

type payload struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only Post requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var payload payload
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		fmt.Printf("could not decode JSON: %s\n", err)
		return
	}

	storage.CreateUniqueUser(payload.Name, payload.Surname)

	fmt.Println("Endpoint hit: create-user")
}
