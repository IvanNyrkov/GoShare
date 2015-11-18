package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

var (
	dbUser     = flag.String("db-user", "admin", "Database username")
	dbPassword = flag.String("db-pass", "admin", "Database password")
	dbName     = flag.String("db-name", "go-share", "Database name")
	db         *sql.DB
)

func init() {
	flag.Parse()
	db = mustConnectDatabase()
}

// Init connection to database
func mustConnectDatabase() (db *sql.DB) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", *dbUser, *dbPassword, *dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}
	return
}

// DBInsertUploadedFileInfo inserts a new record about uploaded file into database
func DBInsertUploadedFileInfo(path string, passphrase string) (err error) {
	stmt, err := db.Prepare("INSERT INTO model.uploaded_file(path, passphrase, created_at) VALUES($1, $2, $3);")
	if err != nil {
		return
	}
	_, err = stmt.Exec(path, passphrase, time.Now())
	if err != nil {
		return
	}
	return
}

// DBFindFilePathByPassphrase returns path to uploaded file
func DBFindFilePathByPassphrase(passphrase string) (path string, err error) {
	err = db.QueryRow("SELECT uf.path FROM model.uploaded_file uf where passphrase=$1", passphrase).Scan(&path)
	if err != nil {
		return
	}
	return
}
