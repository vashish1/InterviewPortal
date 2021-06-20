package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vashish1/InterviewPortal/Backend/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = getInterviewDb()

func GetInterviews() (bool, []models.Interview) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.D{{}}

	var result []models.Interview
	cur, err := db.Find(ctx, filter, options.Find())

	if err != nil {
		fmt.Println("Error while finding User", err)
		return false, []models.Interview{}
	}
	for cur.Next(context.TODO()) {
		var elem models.Interview
		err := cur.Decode(&elem)
		if err != nil {
			fmt.Println("Error while decoding user:", err)
			return false, []models.Interview{}
		}
		result = append(result, elem)
	}
	if err := cur.Err(); err != nil {
		fmt.Println("Error in cursor :", err)
	}

	cur.Close(context.TODO())
	return true, result
}

func CheckAvailability(u models.User, start, end time.Time) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"participants", bson.M{
			"$elemMatch": u,
		}},
		{"starttime", bson.M{
			"$gte": primitive.NewDateTimeFromTime(start),
		}},
		{"endtime", bson.M{
			"$lte": primitive.NewDateTimeFromTime(end),
		}},
	}
	err := db.FindOne(ctx, filter)
	if err.Err() != nil {
		fmt.Println(err.Err().Error())
		return true
	}
	return false
}

func InsertInterviewDetails(data models.Interview) (bool, error) {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := db.InsertOne(ctx, data)
	if err != nil {
		return false, err
	}
	fmt.Println("Inserted an interview data")
	return true, nil
}

func UpdateInterview(data models.Interview) (bool, error) {
	if data.StartTime.Before(time.Now()) && data.EndTime.After(time.Now()) {
		return false, errors.New("Interview cannot be scheduled for past time")
	}
	if len(data.Participants) < 2 {
		return false, errors.New("There must be atleast two participants")
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	filter := bson.D{
		{"id", data.ID},
	}
	update := bson.D{
		{
			"$set", bson.D{
				{"start_time", data.StartTime},
				{"end_time", data.EndTime},
				{"participants", data.Participants},
			},
		},
	}
	result := db.FindOneAndUpdate(ctx, filter, update)
	if result.Err() != nil {
		return false, result.Err()
	}
	return true, nil
}
