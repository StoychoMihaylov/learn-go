package router

import "net/http"

type Router interface {
	GET(uri string, function func(response http.ResponseWriter, request *http.Request))
	POST(uri string, function func(response http.ResponseWriter, request *http.Request))
	SERVE(port string)
}
