package ports

import (
	"context"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
)

// UserService defines the primary port for user operations
type UserService interface {
	Register(ctx context.Context, email, password string) (*domain.UserCreated, error)
	Login(ctx context.Context, email, password string) (*domain.LoginResponse, error)
	GetProfile(ctx context.Context, userID string) (*domain.User, error)
	UpdateProfile(ctx context.Context, userID string, name, surname, biographic, location, website string) error
	GetUsers(ctx context.Context, currentUserID string, page int64, search, userType string) ([]*domain.User, error)
	UploadAvatar(ctx context.Context, userID string, fileData []byte, fileName string) error
	UploadBanner(ctx context.Context, userID string, fileData []byte, fileName string) error
	GetAvatar(ctx context.Context, userID string) ([]byte, error)
	GetBanner(ctx context.Context, userID string) ([]byte, error)
}

// TweetService defines the primary port for tweet operations
type TweetService interface {
	CreateTweet(ctx context.Context, userID, message string) error
	GetTweetsByUser(ctx context.Context, userID string, page int64) ([]*domain.Tweet, error)
	DeleteTweet(ctx context.Context, tweetID, userID string) error
	GetTweetsFromFollowers(ctx context.Context, userID string, page int) ([]*domain.TweetWithFollowers, error)
}

// RelationService defines the primary port for relation operations
type RelationService interface {
	Follow(ctx context.Context, userID, targetUserID string) error
	Unfollow(ctx context.Context, userID, targetUserID string) error
	IsFollowing(ctx context.Context, userID, targetUserID string) (*domain.RelationResponse, error)
}
