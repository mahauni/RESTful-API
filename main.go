package main

import (
	"io"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #1!\n")
}

func homePage2(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello from a HandleFunc #2!\n")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/endpoint", homePage2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
