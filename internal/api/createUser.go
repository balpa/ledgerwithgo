package api

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	response := map[string]interface{}{
		"intValue":    1234,
		"boolValue":   true,
		"stringValue": "hello!",
		"objectValue": map[string]interface{}{
			"arrayValue": []int{1, 2, 3, 4},
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Printf("could not marshal json: %s\n", err)
		return
	}

	fmt.Println(string(jsonResponse))
	fmt.Fprintf(w, string(jsonResponse))
	fmt.Println("Endpoint Hit: createUser")
}
