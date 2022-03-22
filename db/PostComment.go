package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostComment(t models.Comments) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("comments")

	t.CommentID = primitive.NewObjectID()

	reg := bson.M{
		"_id":    t.CommentID,
		"postID": t.PostID,
		"userID": t.UserID,
		"messC":  t.MessC,
		"dateC":  time.Now(),
	}

	results, err := col.InsertOne(ctx, reg)
	if err != nil {
		return "", false, err
	}

	objID := results.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
