package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres driver
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
	err = db.Ping()
	if err != nil {
		return
	}
	dbConnection = &DBConnection{db}
	return
}

// DBInsertUploadedFileInfo inserts a new record about uploaded file into database
func (db *DBConnection) DBInsertUploadedFileInfo(path, passphrase string, createdAt time.Time) (err error) {
	stmt, err := db.Prepare("INSERT INTO model.uploaded_file(path, passphrase, created_at) VALUES($1, $2, $3);")
	if err != nil {
		return
	}
	_, err = stmt.Exec(path, passphrase, createdAt)
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
