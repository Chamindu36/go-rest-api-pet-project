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
	response.Header().Set("Content-Type", "application/json")

	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err2 := response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		if err2 != nil {
			return
		}
		return
	}
	// Return the posts array as a JSON response
	response.WriteHeader(http.StatusOK)
	_, err = response.Write(result)
	if err != nil {
		return
	}
}

func addPost(response http.ResponseWriter, request *http.Request) {
	var post Post
	// Decode the incoming post json and assign to Post struct
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		return
	}

	// Append the post to the posts array
	posts = append(posts, post)

	// Return the posts array as a JSON response
	response.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err2 := response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		if err2 != nil {
			return
		}
		return
	}

	// Set the id of the post
	post.Id = len(posts) + 1

	// Append the post to the posts array
	posts = append(posts, post)

	// Return the posts array as a JSON response
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	_, err = response.Write(result)
	if err != nil {
		return
	}
}
