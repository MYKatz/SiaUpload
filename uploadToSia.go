package main

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"math/rand"
	//"mime/multipart"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UploadedFile struct {
	gorm.Model
	shortName string
	filePath string
}

func addToDB(short string, path string) error{
	db, err := gorm.Open("sqlite3", "files.db")
	if err != nil {
		return err
	}
	defer db.Close()
	file := UploadedFile{shortName: short, filePath: path}
	db.AutoMigrate(&UploadedFile{})
	//create the entry
	db.Create(&file)
	fmt.Println(file.shortName)
	return nil
}

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
		path := "./tmp/" + uuid.New().String() + "___" + header.Filename
		out, err := os.Create(path)
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
		pushToSia(path, w, r)
	} else {
		fmt.Println("Wrong method. Expected POST")
	}

}

func pushToSia(p string, w http.ResponseWriter, r *http.Request) {
	//temporary placement of this function to test some things while I figure out why Siad isn't working.
	addToDB("lima", p)
	//Sia vars
	SiaPass := "passwd"
	SiaPath := "default"
	//http query - specified by Sia Daemon API - default at localhost:9980
	query := "localhost:9980/renter/upload/" + SiaPath + "?source=" + p
	//send file from disk to SIA blockchain
	resp, err := http.Post(query, "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//on success, remove file from tmp folder
	err = os.Remove(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Uploaded file successfully", SiaPass, resp)
}

func generateString() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ$-_!*+(),"
	bytearray := make([]byte, rand.Int())
    for i := range bytearray {
        bytearray[i] = letters[rand.Intn(len(letters))]
    }
    return string(bytearray)
}
