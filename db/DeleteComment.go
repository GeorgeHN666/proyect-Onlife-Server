package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteComment(commentID, userID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("comments")

	comment, _ := primitive.ObjectIDFromHex(commentID)

	condition := bson.M{
		"_id":    comment,
		"userID": userID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil

}
