package server

import (
	"fmt"
	"net/http"
)

/*
net/http is Go’s built-in package to:

-Create web servers
-Handle HTTP requests/responses
-Build APIs

No external libraries needed
*/

// A handler is a function that responds to an HTTP request.

/*

func main() {
	http.HandleFunc("/", server.Home)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

Breakdown of both functions:
w http.ResponseWriter → used to send response
r *http.Request → contains request data
HandleFunc("/", server.Home) → route mapping

*/

/* Sending Responses

Plain text:
fmt.Fprintln(w, "Hello world")

*/

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Go server")

	//Json format response
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message":"hello"}`)

	//Error handling
	http.Error(w, "Something went wrong", http.StatusInternalServerError)

	//Manual Error Handling
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "Status Not found")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Contact Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About Page")
}
