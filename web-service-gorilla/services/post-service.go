package services

import (
	"errors"
	"math/rand"

	entities "web-service-gorilla/entities"
	repositories "web-service-gorilla/repositories/post"
)

type PostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type service struct{}

var (
	repository repositories.PostRepository
)

func NewPostService(repository repositories.PostRepository) PostService {
	repository = repository
	return &service{}
}

func (*service) Validate(post *entities.Post) error {
	if post == nil {
		err := errors.New("The post is empty.")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty.")
		return err
	}

	return nil
}

func (*service) Create(post *entities.Post) (*entities.Post, error) {
	post.ID = rand.Int63()
	return repository.Save(post)
}

func (s *service) FindAll() ([]entities.Post, error) {
	return repository.FindAll()
}
