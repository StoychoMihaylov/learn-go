package router

import (
	"fmt"
	"net/http"

	chi "github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, function func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Get(uri, function)
}

func (*chiRouter) POST(uri string, function func(response http.ResponseWriter, request *http.Request)) {
	chiDispatcher.Post(uri, function)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}
