package main

import (
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

var (
	port      = flag.Int("port", 3010, "Port to serve on")
	webDir    = flag.String("web-dir", "../client/", "Directory with web files")
	uploadDir = flag.String("upload-dir", "./store/", "Directory with uploaded files")
)

// Main function of application
func main() {
	flag.Parse()

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

	path := *uploadDir + handler.Filename
	err = saveFile(path, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	passphrase := GetRandomSentence("_")
	err = DBInsertUploadedFileInfo(path, passphrase)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write([]byte(passphrase))
}

// File saving in upload directory
func saveFile(path string, file multipart.File) (err error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, file)
	return
}
