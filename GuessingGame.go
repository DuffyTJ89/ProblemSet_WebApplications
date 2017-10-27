// web application.
// Author : Thomas Duffy
//Adapted from : https://stackoverflow.com/questions/26559557/how-do-you-serve-a-static-html-file-using-a-go-web-server

package main

import (
	"log"
	"net/http"
	"html/template"
	"math/rand"
	"strconv"
	"time"
)

type message struct{
	Message string 
	GuessUser string
}

//root request
func server(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	}

// /guess request
func guessHandler(w http.ResponseWriter, r *http.Request){

	rand.Seed(time.Now().UTC().UnixNano())

	//check to see if the cookie is already set. If not set new one
	if _, err := r.Cookie("target"); err != nil{
		http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(19)+1)})
	}

	guessUser := r.URL.Query().Get("guess")
	message := &message{Message : "Guess a number 1 to 20 : ", GuessUser: guessUser}

	//send variables to template
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

