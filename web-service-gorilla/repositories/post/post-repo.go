package repositories

import (
	entities "web-service-gorilla/entities"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}
