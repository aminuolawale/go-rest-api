package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	
	"../entity"
	"../service"
	"../errors"
)

type controller struct {}
var postService service.PostService 

type PostController interface {
	GetPosts(res http.ResponseWriter, req *http.Request)
	CreatePost(res http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService)PostController {
	postService = service
	return &controller{}
}

func(*controller) GetPosts(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err !=nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message:"Error getting the posts"})
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}


func(*controller)  CreatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message:"Error processing requeest"})

		return 
	}
	err = postService.Validate(&post)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message:err.Error()})
		return 
	}
	result, err := postService.Create(&post)
	if err!= nil {
		res.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(res).Encode(errors.ServiceError{Message:"error saving the post"})
	}
	fmt.Println("post has been successfully created")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}