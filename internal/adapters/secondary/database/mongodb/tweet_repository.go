package mongodb

import (
	"context"
	"time"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const tweetsCollection = "tweets"

// TweetRepository implements the TweetRepository interface
type TweetRepository struct {
	db *mongo.Database
}

// NewTweetRepository creates a new TweetRepository instance
func NewTweetRepository(conn *Connection) ports.TweetRepository {
	return &TweetRepository{
		db: conn.GetDatabase(databaseName),
	}
}

// Create creates a new tweet in the database
func (r *TweetRepository) Create(ctx context.Context, tweet *domain.Tweet) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(tweetsCollection)

	register := bson.M{
		"userid":  tweet.UserID,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", err
	}

	objID := result.InsertedID.(primitive.ObjectID)
	return objID.String(), nil
}

// FindByUserID finds tweets by user ID with pagination
func (r *TweetRepository) FindByUserID(ctx context.Context, userID string, page int64) ([]*domain.Tweet, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(tweetsCollection)

	findCondition := bson.M{"userid": userID}
	findOptions := options.Find()
	findOptions.SetLimit(20)
	findOptions.SetSort(bson.D{{Key: "date", Value: -1}})
	findOptions.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, findCondition, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*domain.Tweet
	for cursor.Next(ctx) {
		var tweet domain.Tweet
		if err := cursor.Decode(&tweet); err != nil {
			return nil, err
		}
		results = append(results, &tweet)
	}

	return results, cursor.Err()
}

// Delete deletes a tweet
func (r *TweetRepository) Delete(ctx context.Context, tweetID, userID string) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(tweetsCollection)

	objID, err := primitive.ObjectIDFromHex(tweetID)
	if err != nil {
		return err
	}

	condition := bson.M{
		"_id":    objID,
		"userid": userID,
	}

	_, err = col.DeleteOne(ctx, condition)
	return err
}

// FindTweetsFromFollowers finds tweets from followed users
func (r *TweetRepository) FindTweetsFromFollowers(ctx context.Context, userID string, page int) ([]*domain.TweetWithFollowers, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	relationsCol := r.db.Collection("relations")
	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": userID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweets",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "tweet",
		},
	})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := relationsCol.Aggregate(ctx, conditions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*domain.TweetWithFollowers
	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
