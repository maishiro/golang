package entity

import (
	"xorm.io/xorm"

	_ "github.com/lib/pq"
)

var (
	Engine *xorm.Engine
)

func NewConnectDB() error {
	var err error
	Engine, err = xorm.NewEngine("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	return err
}

func CloseDB() {
	Engine.Close()
}
