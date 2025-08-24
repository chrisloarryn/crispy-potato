package mongodb

import (
	"context"
	"time"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const relationsCollection = "relations"

// RelationRepository implements the RelationRepository interface
type RelationRepository struct {
	db *mongo.Database
}

// NewRelationRepository creates a new RelationRepository instance
func NewRelationRepository(conn *Connection) ports.RelationRepository {
	return &RelationRepository{
		db: conn.GetDatabase(databaseName),
	}
}

// Create creates a new relation in the database
func (r *RelationRepository) Create(ctx context.Context, relation *domain.Relation) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(relationsCollection)
	_, err := col.InsertOne(ctx, relation)
	return err
}

// Delete removes a relation from the database
func (r *RelationRepository) Delete(ctx context.Context, relation *domain.Relation) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(relationsCollection)
	_, err := col.DeleteOne(ctx, relation)
	return err
}

// Exists checks if a relation exists in the database
func (r *RelationRepository) Exists(ctx context.Context, relation *domain.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(relationsCollection)

	condition := bson.M{
		"userid":         relation.UserID,
		"userrelationid": relation.UserRelationID,
	}

	var result domain.Relation
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
