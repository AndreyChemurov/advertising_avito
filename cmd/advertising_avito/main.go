package main

import (
	"advertising_avito/internal/restapi"
	"log"
	"net/http"
)

func main() {
	// Routes
	http.HandleFunc("/", restapi.NotFound)
	http.HandleFunc("/create", restapi.Create)
	http.HandleFunc("/get-one", restapi.GetOne)
	http.HandleFunc("/get-all", restapi.GetAll)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
