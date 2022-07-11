package main

import (
	"fmt"
	"log"
	"net/http"

	"mahauni.com/api/routes"
)

func handleRequests() {
	http.HandleFunc("/", routes.HomePage)
	http.HandleFunc("/endpoint", routes.HomePage2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Listening in port: 8080")
	handleRequests()
}
