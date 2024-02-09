package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array."}`))
		return
	}

	response.WriteHeader(http.StatusOK)
	response.Write(result)
	return
}

func addPost(response http.ResponseWriter, request *http.Request) {
	var post Post
	err := json.NewDecoder(request.Body).Decode(&post)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error unmarshalling the request."}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)

	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	response.Write(result)
	return
}
