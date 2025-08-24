package bd

import (
	"context"
	"time"

	"github.com/ccontreras/crispy-potato/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertRegister function inserts a register into the database
func InsertRegister(u models.User) (models.CreatedUser, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	var cu models.CreatedUser

	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return cu, false, err
	}

	cu.ID = result.InsertedID.(primitive.ObjectID)
	cu.Email = u.Email

	return cu, true, nil
}
