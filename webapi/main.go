package main

import (
	"log"
	"net/http"

	"webapi/model"

	sw "webapi/server"

	_ "github.com/lib/pq"
)

func main() {
	err := model.NewConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer model.CloseDB()

	log.Printf("Server start")

	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
