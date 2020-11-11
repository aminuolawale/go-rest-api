package service
import (
	"testing"
	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

)

type MockRepository struct {
	mock.Mock
}

func(mock *MockRepository) Save(post *entity.Post)(*entity.Post, error){
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func(mock *MockRepository) FindAll() ([]entity.Post, error){
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}


func TestFindAll(t *testing.T){
	mockRepo := new(MockRepository)
	var identifier int64 = 1
	post := entity.Post{ID:	identifier, Title:"The title", Text:"Some random stuff"}
	// setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)
	testService := NewPostService(mockRepo)
	result, _ := testService.FindAll()
	// mock assertion: behavioural
	mockRepo.AssertExpectations(t)

	// data assertions
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "The title", result[0].Title)
	assert.Equal(t, "Some random stuff", result[0].Text)


}

func TestValidateEmptyPost(t *testing.T){
	testService := NewPostService(nil)
	err := testService.Validate(nil)
	assert.NotNil(t, err)
	assert.Equal(t, "The post data is empty", err.Error(), )
}


func TestValidateEmptyTitle(t *testing.T){
	testService := NewPostService(nil)
	post := entity.Post{ID:	1, Title:"", Text:"Some random stuff"}
	err := testService.Validate(&post)
	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Error(), )
}


func TestCreate(t *testing.T){
	mockRepo := new(MockRepository)
	post := entity.Post{ Title:"The title", Text:"Some random stuff"}


	// set up expectatios
	mockRepo.On("Save").Return(&post, nil)
	testService := NewPostService(mockRepo)
	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)
	// data 
	assert.NotNil(t, result.ID)
	assert.Equal(t, "The title", result.Title)
	assert.Equal(t, "Some random stuff", result.Text)
	assert.Nil(t, err)

	
}