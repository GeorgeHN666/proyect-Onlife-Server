package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReadOwnPost struct {
	PostID     primitive.ObjectID `bson:"_id"   json:"postID,omitempty"`
	UserIssued primitive.ObjectID `bson:"userIssued"    json:"userIssued,omitempty"`
	MessPost   string             `bson:"messPost"   json:"messPost,omitempty"`
	DatePost   time.Time          `bson:"datePost"   json:"datePost,omitempty"`
}
