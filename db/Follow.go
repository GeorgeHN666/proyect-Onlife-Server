package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Follow(t models.Follow) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("relation")

	t.FollowID = primitive.NewObjectID()

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
