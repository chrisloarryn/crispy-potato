package services

import (
	"context"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// TweetServiceImpl implements the TweetService interface
type TweetServiceImpl struct {
	tweetRepo ports.TweetRepository
}

// NewTweetService creates a new TweetService instance
func NewTweetService(tweetRepo ports.TweetRepository) ports.TweetService {
	return &TweetServiceImpl{
		tweetRepo: tweetRepo,
	}
}

// CreateTweet creates a new tweet
func (s *TweetServiceImpl) CreateTweet(ctx context.Context, userID, message string) error {
	tweet, err := domain.NewTweet(userID, message)
	if err != nil {
		return err
	}

	_, err = s.tweetRepo.Create(ctx, tweet)
	return err
}

// GetTweetsByUser retrieves tweets by user ID
func (s *TweetServiceImpl) GetTweetsByUser(ctx context.Context, userID string, page int64) ([]*domain.Tweet, error) {
	return s.tweetRepo.FindByUserID(ctx, userID, page)
}

// DeleteTweet deletes a tweet
func (s *TweetServiceImpl) DeleteTweet(ctx context.Context, tweetID, userID string) error {
	return s.tweetRepo.Delete(ctx, tweetID, userID)
}

// GetTweetsFromFollowers retrieves tweets from followed users
func (s *TweetServiceImpl) GetTweetsFromFollowers(ctx context.Context, userID string, page int) ([]*domain.TweetWithFollowers, error) {
	return s.tweetRepo.FindTweetsFromFollowers(ctx, userID, page)
}
