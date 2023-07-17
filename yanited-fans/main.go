package main

import (
	"fmt"
	"log"
	"net/http"
)

func fplPathHandler(w http.ResponseWriter, r *http.Request) {
	// Check for correct path entry
	if r.URL.Path != "/fpl" {
		http.Error(w, "404 NOT FOUND, SORRY", http.StatusNotFound)
		return
	}

	// Check for correct method entry
	// Ensure that GET request is handled only
	if r.Method != "GET" {
		http.Error(w, "Method not allowed by server", http.StatusNotFound)
	}

}

func transfersPathHandler(w http.ResponseWriter, r *http.Request) {
	// Check for correct path entry
	if r.URL.Path != "/transfers" {
		http.Error(w, "404 NOT FOUND, SORRY", http.StatusNotFound)
		return
	}

	// Check for correct method entry
	// Ensure that GET request is handled only
	if r.Method != "GET" {
		http.Error(w, "Method not allowed by server", http.StatusNotFound)
	}

}

func joinPathHandler(w http.ResponseWriter, r *http.Request) {
	// Check for errors in the entry
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, ">> ERROR FAIL ~ (POST) /join %v \n", err)
		return
	}
	fmt.Fprintf(w, "SUCESS ~ DATA RECEIVED \n")
	username := r.FormValue("username")
	email := r.FormValue("email")
	fmt.Fprintf(w, "Welcome %s to our fan server.\n", username)
	fmt.Fprintf(w, "An email has been sent to %s.", email)
}

func main() {
	// Define handlers for our requests

	// Check the static directory
	// Handle will use object to handle requests
	//  HandleFunc will use functions to handle requests
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/fpl", fplPathHandler)
	http.HandleFunc("/transfers", transfersPathHandler)
	http.HandleFunc("/join", joinPathHandler)

	// Print Out that server is running
	fmt.Printf(">>> Starting Yanited Fans Server at port 8080/ \n")

	// Check on the errors
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
