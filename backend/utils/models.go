package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	Name string `bson:"name,omitempty" json:"name"`
}

type Shift struct {
	ID primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at"`
	AssignUserId primitive.ObjectID `bson:"assign_user_id,omitempty" json:"assign_user_id"`
	StartDate primitive.DateTime `bson:"start_date,omitempty" json:"start_date"` // default tome RFC3339
	EndDate primitive.DateTime `bson:"end_date,omitempty" json:"end_date"` // default time  RFC3339
	Status string `bson:"status,omitempty" json:"status"`
}
