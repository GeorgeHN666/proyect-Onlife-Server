package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ReadComments(ID string, pg int64) ([]*models.ReadComments, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("comments")

	var results []*models.ReadComments

	condition := bson.M{
		"postID": ID,
	}

	opt := options.Find()

	opt.SetLimit(20)

	opt.SetSort(bson.D{{Key: "dateC", Value: -1}})

	opt.SetSkip((pg - 1) * 20)

	cursor, err := col.Find(ctx, condition, opt)
	if err != nil {
		return results, false
	}

	for cursor.Next(context.TODO()) {

		var registry models.ReadComments

		err := cursor.Decode(&registry)
		if err != nil {
			return results, false
		}

		results = append(results, &registry)
	}

	return results, true

}
