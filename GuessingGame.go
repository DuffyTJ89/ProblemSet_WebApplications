// web application.
// Author : Thomas Duffy
//Adapted from : https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
	"log"
	"net/http"
)


func main() {


	http.Handle("/", http.FileServer(http.Dir("./WebAppsProblemSheet")))
	
	log.Println("Listening....")
    http.ListenAndServe(":8080", nil)
}

