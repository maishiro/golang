package main

import (
	"log"
	"net/http"

	sw "./server"
)

func main() {
	log.Printf("Server start")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
