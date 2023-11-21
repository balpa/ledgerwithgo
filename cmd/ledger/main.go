package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type payload struct {
	Name string `json:"name"`
}

func welcome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprintf(w, "Welcome!!!")
	fmt.Println("Endpoint Hit: welcome")
}

func handleRequests() {
	r := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	http.Handle("/", corsHandler.Handler(r))

	r.HandleFunc("/api/welcome", welcome)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequests()
}
