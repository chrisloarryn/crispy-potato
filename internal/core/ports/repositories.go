package ports

import (
	"context"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.UserCreated, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, bool, error)
	FindByID(ctx context.Context, id string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User, id string) error
	FindAll(ctx context.Context, currentUserID string, page int64, search, userType string) ([]*domain.User, error)
}

// TweetRepository defines the interface for tweet data operations
type TweetRepository interface {
	Create(ctx context.Context, tweet *domain.Tweet) (string, error)
	FindByUserID(ctx context.Context, userID string, page int64) ([]*domain.Tweet, error)
	Delete(ctx context.Context, tweetID, userID string) error
	FindTweetsFromFollowers(ctx context.Context, userID string, page int) ([]*domain.TweetWithFollowers, error)
}

// RelationRepository defines the interface for relation data operations
type RelationRepository interface {
	Create(ctx context.Context, relation *domain.Relation) error
	Delete(ctx context.Context, relation *domain.Relation) error
	Exists(ctx context.Context, relation *domain.Relation) (bool, error)
}

// PasswordHasher defines the interface for password hashing
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

// TokenGenerator defines the interface for JWT token generation
type TokenGenerator interface {
	Generate(user *domain.User) (string, error)
	Validate(token string) (*domain.User, error)
}

// FileStorage defines the interface for file storage operations
type FileStorage interface {
	Save(path string, data []byte) error
	Get(path string) ([]byte, error)
	Delete(path string) error
}
