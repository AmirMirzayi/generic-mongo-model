package db

import (
	"context"
	"fmt"
)

type Subscriber struct {
	Id     string `bson:"_id,omitempty"`
	Name   string `bson:"name,omitempty"`
	Avatar string `bson:"avatar,omitempty"`
}

func (Subscriber) getCollection() string {
	return "subscriber"
}

func InsertSubscriber() {

	newSubscriber := &Subscriber{Name: "8282", Avatar: "https://instagram.com/myavatar"}

	result, err := db.Collection("subscriber").InsertOne(context.TODO(), newSubscriber)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the result is: %v", result)
}
