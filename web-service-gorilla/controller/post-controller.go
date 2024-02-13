package controller

import (
	"encoding/json"
	"net/http"

	entities "web-service-gorilla/entities"
	errors "web-service-gorilla/errors"
	service "web-service-gorilla/service"
)

type controller struct{}

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts."})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	var post entities.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling the request."})
		return
	}

	serviceError := postService.Validate(&post)
	if serviceError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: serviceError.Error()})
		return
	}

	result, serviceError := postService.Create(&post)
	if serviceError != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: serviceError.Error()})
		return
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
