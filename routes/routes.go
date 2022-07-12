package routes

import (
	"log"
	"net/http"

	"mahauni.com/api/services"
)

func HandleRequests() {
	http.HandleFunc("/", services.HomePage)
	http.HandleFunc("/endpoint", services.HomePage2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
