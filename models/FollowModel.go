package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Follow struct {
	FollowID   primitive.ObjectID `bson:"_id"   json:"_id"`
	UserID     string             `bson:"userID"   json:"userID,omitempty"`          // Who I Will Follow
	RelationID string             `bson:"relationID"    json:"relationID,omitempty"` //Who Iam
}
