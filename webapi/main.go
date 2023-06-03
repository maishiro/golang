package main

import (
	"log"

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
	router.Listen(":8080")
}
