package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyProfile(m models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("user")

	reg := make(map[string]interface{})

	if len(m.Name) > 0 {
		reg["name"] = m.Name
	}

	if len(m.Last) > 0 {
		reg["last"] = m.Last
	}

	if len(m.Bio) > 0 {
		reg["bio"] = m.Bio
	}

	if len(m.Avatar) > 0 {
		reg["avatar"] = m.Avatar
	}

	if len(m.Banner) > 0 {
		reg["banner"] = m.Banner
	}

	if len(m.Loc) > 0 {
		reg["loc"] = m.Loc
	}

	if len(m.Web) > 0 {
		reg["web"] = m.Web
	}

	reg["bDay"] = m.Bday

	update := bson.M{
		"$set": reg,
	}

	oID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": oID}}

	_, err := col.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
