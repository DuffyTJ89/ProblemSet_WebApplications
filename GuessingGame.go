// web application.
// Author : Thomas Duffy
//Adapted from : https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
	"log"
	"net/http"
)


func server(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "guess.html")
	}


func main() {
	http.HandleFunc("/", server)//"/" handles any requests and passes to server
	
	log.Println("Listening....")
    http.ListenAndServe(":8080", nil)// serve on port 8080
}

