package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/ccontreras/crispy-potato/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllUsers gets all users
func ReadAllUsers(ID string, page int64, search string, userType string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return results, false
	}

	var foundOne, includeOne bool

	for cursor.Next(ctx) {
		var s models.User
		err := cursor.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}
		var r models.Relation
		r.UserID = ID
		r.UserRelationID = s.ID.Hex()

		includeOne = false
		foundOne, _ = ReadRelation(r)
		if userType == "new" && !foundOne {
			includeOne = true
		}
		if userType == "follow" && foundOne {
			includeOne = true
		}
		if r.UserRelationID == ID {
			includeOne = false
		}
		if includeOne {
			s.Password = ""
			s.Biographic = ""
			s.Website = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""

			results = append(results, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cursor.Close(ctx)
	return results, true
}
