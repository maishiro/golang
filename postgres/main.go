package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// 環境変数読み込み
	cfg := loadEnv()
	// fmt.Println(cfg.ConnectionString())

	// PostgreSQLへ接続
	db, err := sql.Open("postgres", cfg.ConnectionString())
	checkError(err)

	err = db.Ping()
	checkError(err)
	// fmt.Println("Successfully created connection to database")

	// DB読み込み
	sql_statement := "select * from status;"
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	var id int
	var time string
	var data int
	for rows.Next() {
		switch err := rows.Scan(&id, &time, &data); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, time, data)
		default:
			checkError(err)
		}
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type DbConfig struct {
	HostName     string
	PortNo       string
	DbName       string
	UserName     string
	UserPassword string
}

// DB接続文字取得
func (cfg *DbConfig) ConnectionString() string {
	var connectionString string = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", cfg.HostName, cfg.PortNo, cfg.DbName, cfg.UserName, cfg.UserPassword)
	return connectionString
}

// 環境変数読み込み
func loadEnv() DbConfig {
	// .envファイルがあれば取り込む
	godotenv.Load()

	// 環境変数を取得する
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	dbname := os.Getenv("PGDATABASE")
	username := os.Getenv("PGUSER")
	userpass := os.Getenv("PGPASSWORD")
	// fmt.Println(host)
	// fmt.Println(port)
	// fmt.Println(dbname)
	// fmt.Println(username)
	// fmt.Println(userpass)

	var cfg DbConfig
	cfg.HostName = host
	cfg.PortNo = port
	cfg.DbName = dbname
	cfg.UserName = username
	cfg.UserPassword = userpass

	return cfg
}
