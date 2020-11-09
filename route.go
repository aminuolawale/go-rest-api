package main

import (
	"encoding/json"
	"net/http"
	"fmt"
)

type Post struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Text string `json:"text"`
}


var posts []Post

func init(){
	posts = []Post{Post{Id:1, Title:"Title One", Text:"This is the text" }}
}

func getPosts(res http.ResponseWriter, req *http.Request){
	res.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)
	if err !=nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return 
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}


func createPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error processing request"}`))
		return 
	}
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	res.Write(result)
}