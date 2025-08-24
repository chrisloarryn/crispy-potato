package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TweetWithFollowers represents a tweet from followed users
type TweetWithFollowers struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json:"userRelationId,omitempty"`
	Tweet          struct {
		Message string    `bson:"message" json:"message,omitempty"`
		Date    time.Time `bson:"date" json:"date,omitempty"`
		ID      string    `bson:"_id" json:"_id,omitempty"`
	}
}

// RelationResponse represents the status for relations between users
type RelationResponse struct {
	Status bool `json:"status"`
}

// LoginResponse represents the response for login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
