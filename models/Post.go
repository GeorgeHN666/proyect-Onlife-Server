package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	PostID     primitive.ObjectID `bson:"_id"    json:"_id"`
	UserIssued string             `bson:"userIssued"    json:"userIssued"`
	MessPost   string             `bson:"messPost"      json:"messPost"`
	DatePost   time.Time          `bson:"dateTime"      json:"dateTime,omitempty"`
}
