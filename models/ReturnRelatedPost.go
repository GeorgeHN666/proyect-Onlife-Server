package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReturnRelatedPost struct {
	ID         primitive.ObjectID `bson:"_id"    json:"_id,omitempty"`
	UserID     string             `bson:"userID"   json:"userID,omitempty"`
	RelationID string             `bson:"relationID"   json:"relationID,omitempty"`
	Post       struct {
		PostID   string    `bson:"_id"     json:"postID,omitempty"`
		MessPost string    `bson:"messPost"   json:"messPost,omitempty"`
		DatePost time.Time `bson:"datePost"   jjson:"datePost,omitempty"`
	}
}
