package main

import (
	"flag"
	"fmt"
	"github.com/IvanNyrkov/go-share/src/database"
	"github.com/IvanNyrkov/go-share/src/passphrase"
	"github.com/IvanNyrkov/go-share/src/store"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	port   = flag.Int("port", 3010, "Port to serve on")
	webDir = flag.String("web-dir", "../public/", "Directory with web files")

	dbUser     = flag.String("db-user", "admin", "Database username")
	dbPassword = flag.String("db-pass", "admin", "Database password")
	dbName     = flag.String("db-name", "go-share", "Database name")

	db *database.DBConnection
)

// Main function of application
func main() {
	flag.Parse()

	dbConnection, err := database.NewConnection(dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println(err)
		return
	}
	db = dbConnection

	go store.DaemonFileCleaner(*dbConnection)

	router := mux.NewRouter()
	router.HandleFunc("/upload", fileUploadHandler).Methods("POST")

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(*webDir)))
	n.UseHandler(router)

	address := fmt.Sprintf("127.0.0.1:%d", *port)
	n.Run(address)
}

// Handles post-request with file in the body
// Responses with generated passphrase or error
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = store.SaveFile(handler.Filename, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	passphrase := passphrase.GetRandomSentence("_")

	err = db.DBInsertUploadedFileInfo(handler.Filename, passphrase)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(passphrase))
}
