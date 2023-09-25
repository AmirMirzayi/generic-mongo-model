package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Media struct {
	Id      string `bson:"_id,omitempty"`
	Name    string `bson:"name,omitempty"`
	Address string `bson:"address,omitempty"`
}

func (Media) getCollection() string {
	return "media"
}

func InsertMeida() {

	newMedia := &Media{Name: "8282", Address: "https://google.com/"}

	result, err := db.Collection("media").InsertOne(context.TODO(), newMedia)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the result is: %v", result)
}

func FindMedia(id string) *Media {
	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Second)
	defer cancel()

	var result Media
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.M{"_id": objId}
	err = db.Collection("media").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("no document found")
		}
		return nil
	}

	return &result
}
