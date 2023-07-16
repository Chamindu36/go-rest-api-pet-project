package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"golang-rest-api/entity"
	"log"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct{}

const (
	ProjectID           string = "go-rest-api-b1cc2"
	PostsCollectionName        = "posts"
)

// NewPostRepository : New post repository
func NewPostRepository() PostRepository {
	return &repo{}
}

// Save : Implement the Save method of the PostRepository interface
func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()

	// Get a Firestore client
	client, err := firestore.NewClient(ctx, ProjectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
	}

	// Close client when done adding data
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Failed to close Firestore client: %v", err)
		}
	}(client)

	// Add a new document to a collection
	docRef, _, err1 := client.Collection(PostsCollectionName).Add(ctx, map[string]interface{}{
		"Id":    post.Id,
		"Title": post.Title,
		"Text":  post.Text,
	})
	log.Println("Document reference: ", docRef)

	if err1 != nil {
		log.Fatalf("Failed to add a new post: %v", err1)
		return nil, err1
	}

	return post, nil
}

// FindAll : Implement the FindAll method of the PostRepository interface
func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()

	// Get a Firestore client
	client, err := firestore.NewClient(ctx, ProjectID)
	if err != nil {
		log.Fatalf("Failed to create a Firestore client: %v", err)
	}

	// Close client when done adding data
	defer func(client *firestore.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("Failed to close Firestore client: %v", err)
		}
	}(client)

	var posts []entity.Post

	// Get all documents in a collection
	docs, err1 := client.Collection(PostsCollectionName).Documents(ctx).GetAll()
	if err1 != nil {
		log.Fatalf("Failed to get all posts: %v", err1)
		return nil, err1
	}

	// Loop through documents
	for _, doc := range docs {
		post := entity.Post{
			Id:    doc.Data()["Id"].(int64), // Type assertion
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
