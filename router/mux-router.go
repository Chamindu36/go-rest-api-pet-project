package router

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type muxRouter struct{}

var (
	router = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	// Handle get request at /posts
	router.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {

	// Handle post request at /posts
	router.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {

	// Start the server
	log.Println("Mux server is running on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
