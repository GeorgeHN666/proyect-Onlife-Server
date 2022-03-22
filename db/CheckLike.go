package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckLike(t models.Likes) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("likes")

	condition := bson.M{
		"postID": t.PostID,
		"userID": t.UserID,
	}

	var results models.Likes

	err := col.FindOne(ctx, condition).Decode(&results)
	if err != nil {
		return false, err
	}

	return true, nil

}
