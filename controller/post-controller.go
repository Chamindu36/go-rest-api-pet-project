package controller

import (
	"encoding/json"
	"golang-rest-api/entity"
	"golang-rest-api/errors"
	"golang-rest-api/service"
	"math/rand"
	"net/http"
)

var (
	postService service.PostService = service.NewPostService()
)

type controller struct{}

func NewPostController() PostController {
	return &controller{}
}

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

// GetPosts : Get all the posts
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	// Get all the posts from DB
	posts, err0 := postService.FindAll()
	if err0 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts from the database"})
		if err != nil {
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

// AddPost : Add a new post
func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	var post entity.Post

	// Decode the incoming post json and assign to Post struct
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		if err != nil {
			return
		}
		return
	}

	// Set the id of the post
	post.Id = rand.Int63()

	_, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		err3 := json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		if err3 != nil {
			return
		}
		return
	}

	// Validate the post
	err3 := postService.Validate(&post)
	if err3 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		err4 := json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error validating the post"})
		if err4 != nil {
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
