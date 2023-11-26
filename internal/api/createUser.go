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

	// response := map[string]interface{}{
	// 	"intValue":    1234,
	// 	"boolValue":   true,
	// 	"stringValue": "hello!",
	// 	"objectValue": map[string]interface{}{
	// 		"arrayValue": []int{1, 2, 3, 4},
	// 	},
	// }

	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	fmt.Printf("could not marshal json: %s\n", err)
	// 	return
	// }

	// fmt.Fprintf(w, string(jsonResponse))
	fmt.Println("Endpoint hit: create-user")
}
