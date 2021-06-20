package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Createdb creates a database
func Connectdb() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := os.Getenv("DbURL")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		db,
	))
	if err != nil {
		fmt.Println("MongoDB connection failure : ", err)
	}
	return client
}

func getUserDb() *mongo.Collection {
	client := Connectdb()
	userdb := client.Database("InterviewPortal").Collection("User")
	fmt.Println("Connected to User collection")
	return userdb
}

func getInterviewDb() *mongo.Collection {
	client := Connectdb()
	db := client.Database("InterviewPortal").Collection("Interview")
	fmt.Println("Connected to Interview collection")
	return db
}
