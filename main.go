package main

import (
	"fmt"

	"mahauni.com/api/routes"
)

func main() {
	fmt.Println("Listening in port: 8080")
	routes.HandleRequests()
}
