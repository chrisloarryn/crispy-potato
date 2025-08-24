package bd

import (
	"context"
	"log"
	"time"

	"github.com/ccontreras/crispy-potato/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets function for reading tweets
func ReadTweets(ID string, page int64) ([]*models.ReadTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	var results []*models.ReadTweets

	findCondition := bson.M{
		"userid": ID,
	}
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, findCondition, findOptions)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ReadTweets
		err := cursor.Decode(&register)
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}

	return results, true
}
