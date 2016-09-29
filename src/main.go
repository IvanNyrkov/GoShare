package main

import (
	"log"

	"github.com/IvanNyrkov/GoShare/src/rand/sentence"
)

var (
	port   = 3010
	webDir = "public"
)

// Main function of application
func main() {
	sentenceM, err := sentence.NewModule(sentence.ModuleConfig{})
	if err != nil {
		panic(err)
	}
	sentenceS := sentenceM.GetService()
	log.Print(sentenceS.GetRandomSentence("-"))
	log.Print(sentenceS.GetRandomSentence("-"))
	log.Print(sentenceS.GetRandomSentence("-"))
	log.Print(sentenceS.GetRandomSentence("-"))
	log.Print(sentenceS.GetRandomSentence("-"))
	//router := mux.NewRouter()
	//router.HandleFunc("/api/uploadFile", fileUploadHandler).Methods("POST")
	//http.Handle("/api/", router)
	//http.Handle("/", http.FileServer(http.Dir(*webDir)))
	//http.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", *port), nil)
}

// Handles post-request with file in the body
// Responses with generated passphrase or error
//func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
//	r.ParseMultipartForm(0)
//	file, handler, err := r.FormFile("file")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer file.Close()
//
//	passphrase := rand.GetRandomSentence("_")
//	createdAt := time.Now()
//
//	err = store.SaveFile(createdAt.String(), file)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	err = dbConnection.DBInsertUploadedFileInfo(handler.Filename, passphrase, createdAt)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	w.Write([]byte(passphrase))
//}
