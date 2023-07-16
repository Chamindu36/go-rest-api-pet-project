package main

import (
	"encoding/json"
	"golang-rest-api/entity"
	"golang-rest-api/repository"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

// getPosts: Get all the posts
func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Get all the posts from DB
	posts, err0 := repo.FindAll()
	if err0 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err1 := response.Write([]byte(`{"error": "Error getting the posts from the database"}`))
		if err1 != nil {
			return
		}
		return
	}

	// Return the posts array as a JSON response
	response.WriteHeader(http.StatusOK)
	err := json.NewEncoder(response).Encode(posts)
	if err != nil {
		return
	}
}

// addPost: Add a new post
func addPost(response http.ResponseWriter, request *http.Request) {
	var post entity.Post

	// Decode the incoming post json and assign to Post struct
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err1 := response.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		if err1 != nil {
			return
		}
		return
	}

	// Set the id of the post
	post.Id = rand.Int63()

	_, err2 := repo.Save(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		_, err3 := response.Write([]byte(`{"error": "Error saving the post"}`))
		if err3 != nil {
			return
		}
		return
	}

	// Return the posts array as a JSON response
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err4 := json.NewEncoder(response).Encode(post)
	if err4 != nil {
		return
	}
}
