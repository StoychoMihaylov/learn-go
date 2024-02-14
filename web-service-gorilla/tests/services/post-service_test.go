package services

import (
	"testing"

	entities "web-service-gorilla/entities"
	services "web-service-gorilla/services"

	assert "github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entities.Post) (*entities.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entities.Post), entities.Error(1)
}

func (mock *MockRepository) FindAll() ([]entities.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entities.Post), entities.Error(1)
}

func TestFindAll(test *testing.T) {
	mockRepo := new(MockRepository)

	post := entities.Post{
		ID:    1,
		Title: "Some title",
		Text:  "Some text",
	}

	// Setup expectations
	mockRepo.On("FindAll").Return([]entities.Post{post}, nil)

	testService := services.NewPostService(mockRepo)
	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(test)

	// Data Assertion
	assert.Equal(test, 1, result[0].ID)
	assert.Equal(test, "Some title", result[0].Title)
	assert.Equal(test, "Some text", result[0].Text)
}

func TestValidateEmptyPost(test *testing.T) {
	testService := services.NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(test, err)
	assert.Equal(test, "The post is empty.", err.Error())
}

func TestValidateEmptyPostTitle(test *testing.T) {
	testService := services.NewPostService(nil)

	post := entities.Post{
		ID:    1,
		Title: "",
		Text:  "Some text",
	}

	err := testService.Validate(&post)

	assert.NotNil(test, err)
	assert.Equal(test, "The post title is empty.", err.Error())
}
