package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func fetchAllUsers(c *gin.Context) {
	client := GetClient()
	var users []*Users
	collection := client.Database("coba").Collection("users")
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var item Users
		err = cur.Decode(&item)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": http.StatusNotFound, "result": "Data Tidak Ada"})
		}
		users = append(users, &item)
	}
	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "result": users})
}

func main() {
	router := gin.Default()
	v1 := router.Group("/api/user")
	{
		v1.GET("", fetchAllUsers)
	}
	router.Run(":20001")
}
