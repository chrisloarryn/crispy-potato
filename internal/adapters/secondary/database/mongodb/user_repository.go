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

const (
	usersCollection = "users"
	databaseName    = "twittor"
)

// UserRepository implements the UserRepository interface
type UserRepository struct {
	db *mongo.Database
}

// NewUserRepository creates a new UserRepository instance
func NewUserRepository(conn *Connection) ports.UserRepository {
	return &UserRepository{
		db: conn.GetDatabase(databaseName),
	}
}

// Create creates a new user in the database
func (r *UserRepository) Create(ctx context.Context, user *domain.User) (*domain.UserCreated, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(usersCollection)

	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	userCreated := &domain.UserCreated{
		ID:    result.InsertedID.(primitive.ObjectID),
		Email: user.Email,
	}

	return userCreated, nil
}

// FindByEmail finds a user by email
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(usersCollection)
	condition := bson.M{"email": email}

	var user domain.User
	err := col.FindOne(ctx, condition).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &user, false, nil
		}
		return &user, false, err
	}

	return &user, true, nil
}

// FindByID finds a user by ID
func (r *UserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(usersCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	condition := bson.M{"_id": objID}
	var user domain.User

	err = col.FindOne(ctx, condition).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update updates a user in the database
func (r *UserRepository) Update(ctx context.Context, user *domain.User, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(usersCollection)
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	register := make(map[string]interface{})
	if len(user.Name) > 0 {
		register["name"] = user.Name
	}
	if len(user.Surname) > 0 {
		register["surname"] = user.Surname
	}
	register["birthday"] = user.Birthday
	if len(user.Avatar) > 0 {
		register["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		register["banner"] = user.Banner
	}
	if len(user.Biographic) > 0 {
		register["biographic"] = user.Biographic
	}
	if len(user.Location) > 0 {
		register["location"] = user.Location
	}
	if len(user.Website) > 0 {
		register["website"] = user.Website
	}

	updateString := bson.M{"$set": register}
	filter := bson.M{"_id": objID}

	_, err = col.UpdateOne(ctx, filter, updateString)
	return err
}

// FindAll finds all users with pagination and filters
func (r *UserRepository) FindAll(ctx context.Context, currentUserID string, page int64, search, userType string) ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	col := r.db.Collection(usersCollection)

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		// Apply filtering logic based on userType and relations
		// This is simplified - in a real implementation, you might want to
		// handle the relation filtering at the database level for better performance
		if user.ID.Hex() != currentUserID {
			user.Password = ""
			user.Biographic = ""
			user.Website = ""
			user.Location = ""
			user.Banner = ""
			user.Email = ""
			results = append(results, &user)
		}
	}

	return results, cursor.Err()
}
