package userModel

import (
	"time"
	"os"
	"context"
	"log"
	. "schedule-app/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	. "github.com/gobeam/mongo-go-pagination"
)

var collection = "users"

func Add(user User) User{
	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	if insertResult, err := collection.InsertOne(context.TODO(), user); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			user.ID = oid
		}
	} else {
		log.Fatal(err)
	}
	return user
}

func Get(page int64, limit int64, filter primitive.M) ([]User, PaginationData){
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)
	projection := bson.D{
	}
	
	paginatedData, err := New(collection).Limit(limit).Page(page).Sort("created_at", -1).Select(projection).Filter(filter).Find()
	if err != nil {
		panic(err)
	}

	var users []User
	for _, raw := range paginatedData.Data {
		var user *User
		if marshallErr := bson.Unmarshal(raw, &user); marshallErr == nil {
			users = append(users, *user)
		}
	}

	return users, paginatedData.Pagination
}	

func GetByPhone(phone string)(User, error){
	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection(collection)

	var user User
	err := collection.FindOne(context.TODO(), bson.D{{"phone", phone}}).Decode(&user)

	return user, err
}