package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"./repos"
	"./entity"
	"math/rand"
)




var (
	reposit repos.PostRepo = repos.NewRepository()
)
func getPosts(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	posts, err := reposit.FindAll()
	if err !=nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error fetching posts"}`))
		return 
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(posts)
}


func createPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error processing request"}`))
		return 
	}
	post.ID = rand.Int()
	reposit.Save(&post)
	fmt.Println("post has been successfully created")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(post)
}