package repositories

import (
	"context"
	"log"

	entities "web-service-gorilla/entities"

	"cloud.google.com/go/firestore"
)

type repo struct{}

// NewPostRepository
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "pragmatic-reviews"
	collectionName string = "posts"
)

func (*repo) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatal("Failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entities.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)
	if err != nil {
		log.Fatal("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entities.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatal("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entities.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}

	return posts, nil
}
