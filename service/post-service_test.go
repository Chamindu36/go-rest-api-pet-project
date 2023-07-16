package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang-rest-api/entity"
	"testing"
)

// Create a mock struct that implements the PostRepository interface
type MockRepo struct {
	mock.Mock
}

func (mock *MockRepo) Save(post *entity.Post) (*entity.Post, error) {
	// Mock the behavior of the Save() method
	// Return the mocked values
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepo) FindAll() ([]entity.Post, error) {
	// Mock the behavior of the Save() method
	// Return the mocked values
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	// Call the Validate method on the post with nil value
	err := testService.Validate(nil)

	// Assert that the error is not nil and is equal to the error we expect
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("the post is empty"), err)

}

func TestValidateEmptyPostTitle(t *testing.T) {
	var post = entity.Post{
		Id:    1,
		Title: "",
		Text:  "This is a test post",
	}
	testService := NewPostService(nil)
	err := testService.Validate(&post)

	// Assert that the error is not nil and is equal to the error we expect
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("the post title is empty"), err)
}

func TestValidateEmptyPostText(t *testing.T) {
	var post = entity.Post{
		Id:    1,
		Title: "Test Title",
		Text:  "",
	}
	testService := NewPostService(nil)
	err := testService.Validate(&post)

	// Assert that the error is not nil and is equal to the error we expect
	assert.NotNil(t, err)
	assert.Equal(t, errors.New("the post text is empty"), err)
}

func TestFindAllPosts(t *testing.T) {
	mockRepo := new(MockRepo)

	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{
		{Id: 1, Title: "Title 1", Text: "Text 1"},
		{Id: 2, Title: "Title 2", Text: "Text 2"},
	}, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, "Title 1", result[0].Title)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(MockRepo)

	post := entity.Post{Id: 1, Title: "A", Text: "B"}

	// Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data assertions
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "A", result.Title)
	assert.Equal(t, "B", result.Text)
}
