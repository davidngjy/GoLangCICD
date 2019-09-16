package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HealthCheck hit.")
	fmt.Fprintf(w, "Healthy!")
}

func main() {
	fmt.Println("Listening 8080")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", healthCheck)
	log.Fatal(http.ListenAndServe(":8080", router))
}
