package db

import (
	"context"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SeeAllUsers(ID string, t string, srch string, pg int64) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	db := DBConnection.Database("onlife")

	col := db.Collection("user")

	var results []*models.User

	o := options.Find()

	o.SetSkip((pg - 1) * 20)

	o.SetLimit(20)

	q := bson.M{
		"name": bson.M{"$regex": `(?i)` + srch},
	}

	cursor, err := col.Find(ctx, q, o)
	if err != nil {
		return results, false
	}

	var found, include bool

	for cursor.Next(ctx) {

		var s models.User

		err := cursor.Decode(&s)
		if err != nil {
			return results, false
		}

		var r models.Follow

		r.RelationID = ID
		r.UserID = s.ID.Hex()

		include = false

		found, _ = CheckFollow(r)

		if t == "new" && !found {
			include = true
		}

		if t == "follow" && found {
			include = true
		}

		if r.UserID == ID {
			include = false
		}

		if include {
			s.Pass = ""
			s.Bio = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)

		}
	}

	err = cursor.Err()
	if err != nil {
		return results, false
	}

	cursor.Close(ctx)

	return results, true

}
