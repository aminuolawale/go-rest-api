package repos

import (
	"context"
	"log"
	"../entity"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type repo struct {}

// new firestore repository
func NewFirestoreRepository() PostRepo {
	return &repo{}
}

const projectID string= "golang-rest-api-project"
const collectionName string= "posts"


func(*repo) Save(post *entity.Post)(*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalln("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID": post.ID, 
		"Title": post.Title, 
		"Text":post.Text,
	})

	if err!= nil {
		log.Fatalln("Failed to add a new post: %v", err)
		return nil, err
	}
	return post, nil
}

func(*repo) FindAll()([]entity.Post, error){
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalln("Failed to create a firestore client: %v", err)
		return nil, err
	}
	defer client.Close()
	var posts []entity.Post
	it := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err :=it.Next()
		if err == iterator.Done{
			break
		}
		if err!=nil {
			log.Fatalln("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID:doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}