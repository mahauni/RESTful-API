package routes

import (
	"log"
	"net/http"

	"mahauni.com/api/services"
)

func HandleRequests() {
	http.HandleFunc("/", services.HomePage)
	http.HandleFunc("/accounts", services.GetAccounts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
