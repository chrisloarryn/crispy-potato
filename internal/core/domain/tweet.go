package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Tweet represents a tweet entity in the domain
type Tweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"userid" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}

// NewTweet creates a new tweet with validation
func NewTweet(userID, message string) (*Tweet, error) {
	if err := validateTweetMessage(message); err != nil {
		return nil, err
	}
	if userID == "" {
		return nil, errors.New("user ID is required")
	}

	return &Tweet{
		UserID:  userID,
		Message: message,
		Date:    time.Now(),
	}, nil
}

// validateTweetMessage validates tweet message
func validateTweetMessage(message string) error {
	if len(message) == 0 {
		return errors.New("tweet message cannot be empty")
	}
	if len(message) > 280 {
		return errors.New("tweet message cannot exceed 280 characters")
	}
	return nil
}
