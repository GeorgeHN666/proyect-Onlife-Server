package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadAllPost(ID string, pg int64) ([]models.ReturnRelatedPost, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("relation")

	skip := ((pg - 1) * 20)

	con := make([]bson.M, 0)

	con = append(con, bson.M{"$match": bson.M{"relationID": ID}})

	con = append(con, bson.M{
		"$lookup": bson.M{
			"from":         "post",
			"localField":   "userIssued",
			"foreignField": "userID",
			"as":           "post",
		},
	})

	con = append(con, bson.M{"$unwind": "$post"})
	con = append(con, bson.M{"$sort": bson.M{"post.datePost": -1}})
	con = append(con, bson.M{"$skip": skip})
	con = append(con, bson.M{"$limit": 20})

	cursor, _ := col.Aggregate(ctx, con)

	var results []models.ReturnRelatedPost

	err := cursor.All(ctx, &results)
	if err != nil {
		return results, false
	}

	return results, true

}
