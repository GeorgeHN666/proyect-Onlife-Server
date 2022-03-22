package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUserInDB(t models.User) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("user")

	t.Pass, _ = HashingPassword(t.Pass)

	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return "", false, err
	}

	objID := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil

}
