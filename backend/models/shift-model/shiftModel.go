package shiftModel

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

var collection = "shift"

func Add(shift Shift) Shift{
	shift.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	if insertResult, err := collection.InsertOne(context.TODO(), shift); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			shift.ID = oid
		}
	} else {
		log.Fatal(err)
	}
	return shift
}

func Get(page int64, limit int64, filter primitive.M) ([]Shift, PaginationData){
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	projection := bson.D{
	}
	
	paginatedData, err := New(collection).Limit(limit).Page(page).Sort("created_at", -1).Select(projection).Filter(filter).Find()
	if err != nil {
		panic(err)
	}

	var shifts []Shift
	for _, raw := range paginatedData.Data {
		var shift *Shift
		if marshallErr := bson.Unmarshal(raw, &shift); marshallErr == nil {
			shifts = append(shifts, *shift)
		}
	}

	return shifts, paginatedData.Pagination
}

func Update(ID primitive.ObjectID, shift Shift ) (*mongo.UpdateResult,  error) {
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	log.Println(shift.EndDate)
	result , err := collection.UpdateOne(context.TODO(), bson.M{ "_id" : ID,}, bson.M{
		"$set" : bson.M{
			"start_date" : shift.StartDate,
			"end_date" : shift.EndDate,
			"assign_user_id" : shift.AssignUserId,
		},
	},)

	return result, err
}
