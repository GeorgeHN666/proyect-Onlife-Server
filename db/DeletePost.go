package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeletePost(PostID, UserID string) error {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("post")

	objID, _ := primitive.ObjectIDFromHex(PostID)

	condition := bson.M{
		"_id":        objID,
		"userIssued": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return err
	}

	return nil

}
