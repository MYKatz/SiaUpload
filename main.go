package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("yo!")
		fmt.Fprintf(w, "Hello world!")
	})

	r.HandleFunc("/sendFile", func(w http.ResponseWriter, r *http.Request) {
		//upload to SIA
	})

	http.ListenAndServe(":80", r)
}