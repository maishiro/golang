package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// 引数の数を確認する。
	if len(os.Args) < 3 {
		showCommandArgs()
		os.Exit(1)
	}
	// os.Argsを確認する。
	// fmt.Println(os.Args)

	// コマンド引数解析
	// flag.Parse()
	// fmt.Println(flag.Args())

	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])

	// ファイル読み込み
	sql_statement := ReadSQLFile(os.Args[2])

	// 環境変数読み込み
	cfg := loadEnv()
	// fmt.Println(cfg.ConnectionString())

	// PostgreSQLへ接続
	db, err := sql.Open("postgres", cfg.ConnectionString())
	checkError(err)

	err = db.Ping()
	checkError(err)
	// fmt.Println("Successfully created connection to database")

	switch os.Args[1] {
	case "query":
		doQuery(db, sql_statement)
	case "command":
		doCommand(db, sql_statement)
	default:
		showCommandArgs()
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func showCommandArgs() {
	fmt.Println("ERROR: Command args")
	fmt.Println("  cmd [CommandType command/query] [SQL file path]")
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

// SQLファイルの読み込み
func ReadSQLFile(fileName string) string {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(bytes))
	return string(bytes)
}

// DBクエリー(値が返るとき)
func doQuery(db *sql.DB, sql_statement string) {
	// DB読み込み
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	// // Data Column Type
	// colTypes, err := rows.ColumnTypes()
	// for _, t := range colTypes {
	// 	fmt.Println(t.ScanType())
	// }

	// Data Column Header
	cols, err := rows.Columns()
	fmt.Println(strings.Join(cols, ","))

	// Data
	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}

		err := rows.Scan(rowp...)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			rowValue := []string{}
			for i, _ := range cols {
				switch row[i].(type) {
				case int64:
					row[i] = row[i].(int64)
					rowValue = append(rowValue, fmt.Sprintf("%d", row[i]))
				case time.Time:
					var dt time.Time
					dt = row[i].(time.Time)
					row[i] = dt.Format("2006-01-02 15:04:05-0700")
					rowValue = append(rowValue, dt.Format("2006-01-02 15:04:05-0700"))
				default:
					fmt.Println(row[i])
					fmt.Println(row[i].(string))
					rowValue = append(rowValue, row[i].(string))
				}
			}
			fmt.Println(strings.Join(rowValue, ","))
		default:
			checkError(err)
		}
	}
}

// DB更新コマンド(Insert,Delete,Update,...)
func doCommand(db *sql.DB, sql_statement string) {
	// DB更新
	_, err := db.Exec(sql_statement)
	checkError(err)
}
