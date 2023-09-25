package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const uri = "mongodb://root:root_password@127.0.0.1:27017"

var (
	db     *mongo.Database
	client *mongo.Client
)

type cRUD string

var (
	insert cRUD = "created"
	delete cRUD = "deleted"
	update cRUD = "updated"
	find   cRUD = "founded"
)

type dbModel interface {
	getCollection() string
	Post | Media | Subscriber
}

func Init() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	var err error
	// Create a new client and connect to the server
	client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	db = client.Database("test_generic_model")
}

func Ping() {
	var result bson.M
	if err := db.RunCommand(context.TODO(), bson.D{{Name: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func Close() {
	if err := client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}

func Insert[T dbModel](model T) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	_, err := db.Collection(model.getCollection()).InsertOne(ctx, model)
	if err != nil {
		return err
	}
	go publishAck(model, insert)
	return nil
}

func Find[T dbModel](model T, id string) *T {
	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Second)
	defer cancel()
	var result T
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil
	}
	filter := bson.M{"_id": objId}
	err = db.Collection(model.getCollection()).FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			println("no document found")
		}
		return nil
	}
	go publishAck(model, find)
	return &result
}

func publishAck[T dbModel, V cRUD](model T, method cRUD) {
	fmt.Println("i have " + method.getMethod() + " 1 " + model.getCollection())
}

func (c cRUD) getMethod() string {
	return string(c)
}
