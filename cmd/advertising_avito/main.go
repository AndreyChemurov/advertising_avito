package main

import (
	"advertising_avito/internal/database"
	"advertising_avito/internal/restapi"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Create database and indecies
	if err := database.CreateTableAndIndecies(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Routes
	http.HandleFunc("/", restapi.NotFound)
	http.HandleFunc("/create", restapi.Create)
	http.HandleFunc("/get-one", restapi.GetOne)
	http.HandleFunc("/get-all", restapi.GetAll)

	log.Println("Server is ready to accept requests")

	log.Fatal(http.ListenAndServe(":8000", nil))
}
