package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteFollow(t models.Follow) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("relation")

	condition := bson.M{
		"userID":     t.UserID,
		"relationID": t.RelationID,
	}

	_, err := col.DeleteOne(ctx, condition)
	if err != nil {
		return false, err
	}

	return true, nil

}
