package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadComments struct {
	CommentID primitive.ObjectID `bson:"_id"   json:"_id"`
	PostID    string             `bson:"postID"   json:"postID,omitempty"`
	UserID    string             `bson:"userID"   json:"userID,omitempty"`
	MessC     string             `bson:"messC"    json:"messC,omitempty"`
	DateC     time.Time          `bson:"dateC"    json:"dateC,omitempty"`
}
