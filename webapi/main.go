package main

import (
	"log"
	"net/http"

	"webapi/entity"

	sw "webapi/server"

	_ "github.com/lib/pq"
)

func main() {
	err := entity.NewConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer entity.CloseDB()

	log.Printf("Server start")

	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
