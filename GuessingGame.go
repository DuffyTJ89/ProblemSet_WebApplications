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
	GuessUser int
	Winner bool
}

//root request
func server(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	}

// /guess request
func guessHandler(w http.ResponseWriter, r *http.Request){

	rand.Seed(time.Now().UTC().UnixNano())

	//check to see if the cookie is already set. If not set new one
	cookie, err := r.Cookie("target")
	
	if err != nil {
		http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(19) + 1)})
	}

	//get the users guess from the url query
	guessUser, _ := strconv.Atoi(r.URL.Query().Get("guess"))

	message := &message{Message : "Guess a number 1 to 20 : ", GuessUser: guessUser, Winner: false}

	//convert target to number to compare
	target, _ := strconv.Atoi(cookie.Value)

	//check cookie with guessUser to see if they match
	if (target == guessUser){
		message.Winner = true
		message.Message = "A winner is you"
		http.SetCookie(w, &http.Cookie{Name: "target", Value: strconv.Itoa(rand.Intn(19) + 1)})
	}else if (target < guessUser){
		message.Message = "Too high, guess again.."
	}else if (target > guessUser){
		message.Message = "Too low, guess again.."
	}else {
		message.Message = "Something went wrong..."
	}

	//parse template
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

