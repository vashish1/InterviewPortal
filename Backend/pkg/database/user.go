package database

import (
	"context"
	"fmt"
	"time"

	"github.com/vashish1/InterviewPortal/Backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userdb = getUserDb()


func InsertUserData(userData models.User) (bool, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := userdb.InsertOne(ctx, userData)
	if err != nil {
		return false, err
	}
	fmt.Println("Inserted a user data")
	return true, err
}

func GetUserData() (bool, []models.User) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{}}

	var result []models.User
	cur, err := userdb.Find(ctx, filter, options.Find())

	if err != nil {
		fmt.Println("Error while finding User", err)
		return false, []models.User{}
	}
	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error while decoding user:", err)
			return false, []models.User{}
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("Error in cursor :", err)
	}

	cur.Close(context.TODO())
	return true, result
}
