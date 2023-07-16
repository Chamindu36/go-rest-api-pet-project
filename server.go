package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Create a new router
	var router = mux.NewRouter()

	// Define the port
	const port = ":8000"

	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and Running...")
	})

	// Handle get request at /posts
	router.HandleFunc("/posts", getPosts).Methods("GET")
	// Handle post request at /posts
	router.HandleFunc("/posts", addPost).Methods("POST")

	// Start the server
	log.Println("Server is running on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
