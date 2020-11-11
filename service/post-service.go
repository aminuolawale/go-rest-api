package service

import (
	"../repos"
	"errors"
	"../entity"
	"math/rand"

)

var (
	reposit repos.PostRepo 
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll()([]entity.Post, error)
}


type service struct {}

func NewPostService(repo repos.PostRepo) PostService {
	reposit = repo
	return &service{}
}



func(*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("The post data is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post data is empty")
		return err
	}
	return nil
}


func(*service) Create(post *entity.Post)(*entity.Post, error) {
	post.ID = rand.Int63()
	post, err := reposit.Save(post)
	return post, err
}

func(*service) FindAll() ([]entity.Post, error) {
	return reposit.FindAll()
}