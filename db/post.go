package db

import (
	"context"
)

type Post struct {
	Id      string `bson:"_id,omitempty"`
	Name    string `bson:"name,omitempty"`
	Address string `bson:"address,omitempty"`
}

func (Post) getCollection() string {
	return "post"
}

func InsertPost() {
	newPost := &Post{Name: "8282", Address: "https://google.com/"}

	_, err := db.Collection("post").InsertOne(context.TODO(), newPost)
	if err != nil {
		panic(err)
	}
}
