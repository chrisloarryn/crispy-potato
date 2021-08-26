package models

// Tweet gets from the body the incoming message
type Tweet struct {
	Message string `bson:"message" json:"message"`
}
