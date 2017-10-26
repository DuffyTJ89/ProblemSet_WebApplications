// web application.
// Author : Thomas Duffy
//Adapted from : https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
	"log"
	"net/http"
	"html/template"
)

type message struct{
	Message string 
}

//root request
func server(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	}

// /guess request
func guessHandler(w http.ResponseWriter, r *http.Request){

	message := &message{Message : "Guess a number 1 to 20 : "}

	t, _ := template.ParseFiles("guess.tmpl")
	t.Execute(w, message)

	//http.ServeFile(w, r, "guess.html")
}



func main() {
	http.HandleFunc("/", server)//"/" handles any requests and passes to server
	http.HandleFunc("/guess", guessHandler) //handles requests for guess

	log.Println("Listening....")
    http.ListenAndServe(":8080", nil)// serve on port 8080
}

