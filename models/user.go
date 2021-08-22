package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// User data structure for save the user information
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Surname    string             `bson:"surname" json:"surname,omitempty"`
	Birthday   time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password,omitempty"`
	Avatar     string             `bson:"avatar" json:"avatar,omitempty"`
	Banner     string             `bson:"banner" json:"banner,omitempty"`
	Biographic string             `bson:"biographic" json:"biographic,omitempty"`
	Location   string             `bson:"location" json:"location,omitempty"`
	Website    string             `bson:"website" json:"website,omitempty"`
}
