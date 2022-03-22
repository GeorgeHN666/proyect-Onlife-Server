package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Post(t models.Post) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("post")

	reg := bson.M{
		"userIssued": t.UserIssued,
		"messPost":   t.MessPost,
		"datePost":   t.DatePost,
	}

	results, err := col.InsertOne(ctx, reg)
	if err != nil {
		return "", false, err
	}

	objID := results.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
