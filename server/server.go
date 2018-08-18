package main

import (
	"log"
	"net/http"
)

func main() {
	router := createRouter()
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", router))
}
