package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

type payload struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Token   string `json:"token"`
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

	encodedString := base64.StdEncoding.EncodeToString([]byte(payload.Token))

	storage.CreateUniqueUser(payload.Name, payload.Surname, encodedString)

	fmt.Println("Endpoint hit: create-user")
}
