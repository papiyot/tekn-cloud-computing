package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Users struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func ReturnAllUsers(client *mongo.Client, filter bson.M) []*Users {
	var users []*Users
	collection := client.Database("coba").Collection("users")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var item Users
		err = cur.Decode(&item)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		users = append(users, &item)
	}
	return users
}

func main() {
	c := GetClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	users := ReturnAllUsers(c, bson.M{})
	for _, item := range users {
		log.Println(item.Name, item.Email)
	}
}
