package main

import (
	"flag"
	"fmt"
	"github.com/IvanNyrkov/Go-Share/src/database"
	"github.com/IvanNyrkov/Go-Share/src/passphrase"
	"github.com/IvanNyrkov/Go-Share/src/store"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	port   = flag.Int("port", 3010, "Port to serve on")
	webDir = flag.String("web-dir", "public", "Directory with web files")

	dbUser     = flag.String("db-user", "admin", "Database username")
	dbPassword = flag.String("db-pass", "admin", "Database password")
	dbName     = flag.String("db-name", "go-share", "Database name")

	dbConnection *database.DBConnection
)

// Main function of application
func main() {
	flag.Parse()

	var err error
	dbConnection, err = database.NewConnection(dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println(err)
		return
	}

	go store.DaemonFileCleaner(*dbConnection)

	router := mux.NewRouter()
	router.HandleFunc("/api/uploadFile", fileUploadHandler).Methods("POST")
	http.Handle("/api/", router)
	http.Handle("/", http.FileServer(http.Dir(*webDir)))
	http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), nil)
}

// Handles post-request with file in the body
// Responses with generated passphrase or error
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	passphrase := passphrase.GetRandomSentence("_")
	createdAt := time.Now()

	err = store.SaveFile(createdAt.String(), file)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dbConnection.DBInsertUploadedFileInfo(handler.Filename, passphrase, createdAt)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(passphrase))
}
