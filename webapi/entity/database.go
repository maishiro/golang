package entity

import (
	"database/sql"

	"xorm.io/xorm"
	"xorm.io/xorm/core"

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

// for UnitTest
func NewConnectDBWithDB(db *sql.DB) error {
	var err error
	Engine, err = xorm.NewEngineWithDB("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable", core.FromDB(db))
	return err
}

func CloseDB() {
	Engine.Close()
}
