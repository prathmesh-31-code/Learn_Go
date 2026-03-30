package main

import (
	"fmt"
	"learn_go/server"
	"net/http"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.HandleFunc("/Contact", server.Contact)
	http.HandleFunc("/About", server.About)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
