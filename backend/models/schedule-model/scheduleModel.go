package scheduleModel

import (
	"time"
	"os"
	"context"
	"log"
	. "schedule-app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = "schedule"

func Add(schedule Schedule) Schedule{
	schedule.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	if insertResult, err := collection.InsertOne(context.TODO(), schedule); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			schedule.ID = oid
		}
	} else {
		log.Fatal(err)
	}
	return schedule
}

func Get(page int64, limit int64, filter primitive.M) ([]Schedule, PaginationData){
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	projection := bson.D{
	}
	
	paginatedData, err := New(collection).Limit(limit).Page(page).Sort("created_at", -1).Select(projection).Filter(filter).Find()
	if err != nil {
		panic(err)
	}

	var schedules []Schedule
	for _, raw := range paginatedData.Data {
		var schedule *Schedule
		if marshallErr := bson.Unmarshal(raw, &schedule); marshallErr == nil {
			schedules = append(schedules, *schedule)
		}
	}

	return schedules, paginatedData.Pagination
}

func Update(ID primitive.ObjectID, schedule Schedule ) (*mongo.UpdateResult,  error) {
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	log.Println(schedule.EndDate)
	result , err := collection.UpdateOne(context.TODO(), bson.M{ "_id" : ID,}, bson.M{
		"$set" : bson.M{
			"start_date" : schedule.StartDate,
			"end_date" : schedule.EndDate,
			"assign_user_id" : schedule.AssignUserId,
		},
	},)

	return result, err
}
