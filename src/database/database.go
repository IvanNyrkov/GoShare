package database

import (
	"database/sql"
	"fmt"
	"time"
)

// DBConnection holds connection to database
type DBConnection struct {
	*sql.DB
}

// NewConnection creates connection to database
func NewConnection(dbUser, dbPassword, dbName *string) (dbConnection *DBConnection, err error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", *dbUser, *dbPassword, *dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		return
	}
	return &DBConnection{db}, nil
}

// DBInsertUploadedFileInfo inserts a new record about uploaded file into database
func (db *DBConnection) DBInsertUploadedFileInfo(path string, passphrase string) (err error) {
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
func (db *DBConnection) DBFindFilePathByPassphrase(passphrase string) (path string, err error) {
	err = db.QueryRow("SELECT uf.path FROM model.uploaded_file uf where passphrase=$1", passphrase).Scan(&path)
	if err != nil {
		return
	}
	return
}
