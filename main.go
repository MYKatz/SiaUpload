package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./web/assets/"))))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//send homepage
		tmp, err := template.ParseFiles("web/index.html")
		data := ""

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
        	return
		}

		if err := tmp.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Printf("Served homepage")
	})

	r.HandleFunc("/sendFile", upload)

	http.ListenAndServe(":80", r)
}