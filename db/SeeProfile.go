package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SeeProfile(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("user")

	id, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{"_id": id}

	var results models.User

	err := col.FindOne(ctx, condition).Decode(&results)

	results.Pass = ""

	if err != nil {
		return results, err
	}

	return results, nil

}
