package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct{}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, function func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, function).Methods("GET")
}

func (*muxRouter) POST(uri string, function func(response http.ResponseWriter, request *http.Request)) {
	muxDispatcher.HandleFunc(uri, function).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}
