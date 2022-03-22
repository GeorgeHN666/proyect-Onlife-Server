package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"    json:"_id"`
	Name   string             `bson:"name"   json:"name,omitempty"`
	Last   string             `bson:"last"   json:"last,omitempty"`
	Email  string             `bson:"email"  json:"email"`
	Pass   string             `bson:"pass"   json:"pass,omitempty"`
	Bday   time.Time          `bson:"bDay"   json:"bDay,omitempty"`
	Avatar string             `bson:"avatar" json:"avatar,omitempty"`
	Banner string             `bson:"banner" json:"banner,omitempty"`
	Loc    string             `bson:"loc"    json:"loc,omitempty"`
	Web    string             `bson:"web"    json:"web,omitempty"`
	Bio    string             `bson:"bio"    json:"bio,omitempty"`
}
