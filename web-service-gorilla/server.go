package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8000"
	router := mux.NewRouter()
	router.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Up and running...")
	})

	//APIs
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPost).Methods("POST")

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
