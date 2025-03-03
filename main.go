package main

import (
	"fmt"
	"log"
	"net/http"
)

var Articles []Article

func main() {
	err := InitDB()
	if err != nil {
		log.Fatal(err)
	}
	r := SetupRoutes()
	fmt.Printf("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
