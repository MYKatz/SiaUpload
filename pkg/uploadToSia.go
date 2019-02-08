package uploadToSia

import (
	"fmt"
	"net/http"

)

func upload(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20) //Parses data, w/ filesize limit of 32MiB
		file, handler, err := r.FormFile("fileUpload")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//pushToSia(file)
	} else {
		fmt.Println("Wrong method. Expected POST")
	}

}

func pushToSia(f http.File) {
	fmt.Println("pushing file...")
	//send file to SIA blockchain
}
