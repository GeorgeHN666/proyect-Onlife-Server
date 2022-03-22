package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var POSTID string

func ReadOwnPost(ID string, pg int64) ([]*models.ReadOwnPost, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("post")

	var results []*models.ReadOwnPost

	condition := bson.M{
		"userIssued": ID,
	}

	opt := options.Find()

	opt.SetLimit(20)

	opt.SetSort(bson.D{{Key: "datePost", Value: -1}})

	opt.SetSkip((pg - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		return results, false
	}
	var registry models.ReadOwnPost

	for cursor.Next(context.TODO()) {

		err := cursor.Decode(&registry)
		POSTID = registry.PostID.String()
		if err != nil {
			return results, false
		}

		results = append(results, &registry)
	}

	return results, true

}
