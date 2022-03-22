package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CheckFollow(t models.Follow) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("relation")

	condition := bson.M{
		"userID":     t.UserID,
		"relationID": t.RelationID,
	}

	var results models.CheckRelation

	err := col.FindOne(ctx, condition).Decode(&results)
	if err != nil {
		return false, err
	}

	return true, nil

}
