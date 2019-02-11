package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"mime/multipart"
	"github.com/google/uuid"
)

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20) //Parses data, w/ filesize limit of 32MiB
		file, header, err := r.FormFile("fileUpload")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		//writes file to UUID + ___ + filename in tmp folder
		out, err := os.Create("./tmp/" + uuid.New().String() + "___" + header.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "File uploaded successfully")
		//pushToSia(file)
	} else {
		fmt.Println("Wrong method. Expected POST")
	}

}

func pushToSia(f multipart.File) {
	fmt.Println("pushing file...")
	//send file to SIA blockchain
}
