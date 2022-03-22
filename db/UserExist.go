package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func UserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("user")

	condition := bson.M{"email": email}

	var results models.User

	err := col.FindOne(ctx, condition).Decode(&results)

	id := results.ID.Hex()

	if err != nil {
		return results, false, id
	}

	return results, true, id

}
