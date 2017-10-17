// An  web application.
// Author : Thomas Duffy
//Adapted from : http://blog.scottlogic.com/2017/02/28/building-a-web-app-with-go.html
package main

import (
	"fmt"
	"net/http"
    
)

//function that prints the name to the web page
func printName(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Guessing Game")
}//end printName

func main() {
	//handle requests by calling printName
    http.HandleFunc("/", printName)
	//start webserver and serve on port 8080
    http.ListenAndServe(":8080", nil)
}

