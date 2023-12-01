package api

import (
	"fmt"
	"net/http"

	"github/balpa/ledgerwithgo/internal/storage"
)

func GetAllBalances(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	allBalances := storage.GetAllBalances()

	fmt.Fprintf(w, string(allBalances))
	fmt.Println("Endpoint hit: get-all-balances")
}
