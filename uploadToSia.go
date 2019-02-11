package main

import (
	"fmt"
	"net/http"
	"mime/multipart"
)

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20) //Parses data, w/ filesize limit of 32MiB
		file, _, err := r.FormFile("fileUpload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "Got file.")
		pushToSia(file)
	} else {
		fmt.Println("Wrong method. Expected POST")
	}

}

func pushToSia(f multipart.File) {
	fmt.Println("pushing file...")
	//send file to SIA blockchain
}
