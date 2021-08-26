package bd

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// DeleteTweet is for deleting a tweet
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15 * time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweets")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}
