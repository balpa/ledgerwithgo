package main

import (
	"github/balpa/ledgerwithgo/internal/controllers"
	"github/balpa/ledgerwithgo/internal/storage"
)

func main() {
	storage.ConnectDB()

	controllers.HandleRequests()
}
