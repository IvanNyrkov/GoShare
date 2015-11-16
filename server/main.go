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
	port       = *flag.Int("port", 3010, "Port to serve on")
	webDir     = *flag.String("webDir", "../client/", "Directory with web files")
	uploadFile = *flag.String("uploadFile", "file", "Name of upload file form parameter")
	uploadDir  = *flag.String("uploadDir", "./store/", "Directory with uploaded files")
)

// Main function of application
func main() {
	flag.Parse()

	router := mux.NewRouter()
	router.HandleFunc("/upload", fileUploadHandler).Methods("POST")

	n := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger(),
		negroni.NewStatic(http.Dir(webDir)))
	n.UseHandler(router)

	address := fmt.Sprintf("127.0.0.1:%d", port)
	n.Run(address)
}

// Handles post-request with file in the body
// Responses with generated codephrase or error
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(0)
	file, handler, err := r.FormFile(uploadFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	err = saveFile(handler.Filename, file)
	if err != nil {
		fmt.Println(err)
		return
	}

	codephrase := GetRandomSentence("_")
	w.Write([]byte(codephrase))
}

// File saving in upload directory
func saveFile(fileName string, file multipart.File) (err error) {
	f, err := os.OpenFile(uploadDir+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()

	io.Copy(f, file)
	return
}
