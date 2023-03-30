package db

import (
	"os"
	"log"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_URI"))
	// You need to import `go-sql-driver/mysql` package
	// if you are using "mysql" as first argument to `sql.Open`.
	if err != nil {
		log.Fatal(err)
	}
	return db
}
