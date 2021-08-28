package bd

import (
	"context"
	"github.com/ccontreras/crispy-potato/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

// ReadRelation find the relations
func ReadRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relations")

	condition := bson.M{
		"userid":         t.UserID,
		"userrelationid": t.UserRelationID,
	}

	var result models.Relation

	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		return false, err
	}
	return true, err
}
