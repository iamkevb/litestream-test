package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var createTableName string
var writeDB *sql.DB
var readDB *sql.DB

func init() {
	flag.StringVar(&createTableName, "create", "", "the name of the table to be created")
}

func main() {
	flag.Parse()

	writeDB = openWriteDB()
	defer writeDB.Close()

	readDB = openReadDB()
	defer readDB.Close()

	if len(createTableName) != 0 {
		createTable()
	}
}

func createTable() {
	sql := "CREATE TABLE %s (id INTEGER PRIMARY KEY AUTOINCREMENT, first TEXT, last TEXT);"
	createStmt := fmt.Sprintf(sql, createTableName)
	_, err := writeDB.Exec(createStmt)
	if err != nil {
		log.Fatal(err)
	}

	sql = "SELECT id from %s"
	queryStmt := fmt.Sprintf(sql, createTableName)
	start := time.Now().UnixMilli()

	for {
		_, err := readDB.Query(queryStmt)
		d := time.Now().UnixMilli() - start
		if err != nil {
			time.Sleep(1)
			if d > 1000 {
				log.Fatal("table not replicated after 1 second")
			}
		}
		log.Printf("table %s replicated to read database %dms\n", createTableName, d)
		break
	}
}

func openWriteDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/write.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func openReadDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/read.db")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
