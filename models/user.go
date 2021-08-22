package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id, omitempty" json:"id"`
	Name       string             `bson:"name" json:"name, omitempty"`
	Surname    string             `bson:"surname" json:"name, omitempty"`
	Birthday   time.Time             `bson:"birthday" json:"birthday, omitempty"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"name" json:"name, omitempty"`
	Avatar     string             `bson:"name" json:"name, omitempty"`
	Banner     string             `bson:"name" json:"name, omitempty"`
	Biographic string             `bson:"name" json:"name, omitempty"`
	Location   string             `bson:"location" json:"name, omitempty"`
	Website    string             `bson:"website" json:"name, omitempty"`
}
