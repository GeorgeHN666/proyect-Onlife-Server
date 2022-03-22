package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Likes struct {
	LikeID primitive.ObjectID `bson:"_id"    json:"_id"`
	PostID string             `bson:"postID"   json:"postID,omitempty"`
	UserID string             `bson:"userID"   json:"userID,omitempty"`
}
