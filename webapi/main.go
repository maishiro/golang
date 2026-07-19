package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"webapi/entity"

	sw "webapi/server"

	_ "github.com/lib/pq"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	fileLogger := &lumberjack.Logger{
		Filename:   "./logs/webapi.log",
		MaxSize:    10,
		MaxBackups: 9,
		LocalTime:  true,
		Compress:   false,
	}
	multiWriter := io.MultiWriter(os.Stdout, fileLogger)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	err := entity.NewConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer entity.CloseDB()

	log.Printf("Server start")

	router := sw.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
