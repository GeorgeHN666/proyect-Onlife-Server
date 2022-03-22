package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteLike(likeID, userID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("likes")

	like, _ := primitive.ObjectIDFromHex(likeID)

	condition := bson.M{
		"_id":    like,
		"userID": userID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil

}
