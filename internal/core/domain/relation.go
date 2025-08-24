package domain

import "errors"

// Relation represents a relationship between users
type Relation struct {
	UserID         string `bson:"userid" json:"userId"`
	UserRelationID string `bson:"userrelationid" json:"userRelationId"`
}

// NewRelation creates a new relation with validation
func NewRelation(userID, userRelationID string) (*Relation, error) {
	if userID == "" {
		return nil, errors.New("user ID is required")
	}
	if userRelationID == "" {
		return nil, errors.New("user relation ID is required")
	}
	if userID == userRelationID {
		return nil, errors.New("user cannot follow themselves")
	}

	return &Relation{
		UserID:         userID,
		UserRelationID: userRelationID,
	}, nil
}
