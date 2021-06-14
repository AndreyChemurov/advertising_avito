package main

import (
	"advertising_avito/internal/restapi"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", restapi.NotFound)
	http.HandleFunc("/create", restapi.Create)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
