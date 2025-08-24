package services

import (
	"context"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// RelationServiceImpl implements the RelationService interface
type RelationServiceImpl struct {
	relationRepo ports.RelationRepository
}

// NewRelationService creates a new RelationService instance
func NewRelationService(relationRepo ports.RelationRepository) ports.RelationService {
	return &RelationServiceImpl{
		relationRepo: relationRepo,
	}
}

// Follow creates a new relation between users
func (s *RelationServiceImpl) Follow(ctx context.Context, userID, targetUserID string) error {
	relation, err := domain.NewRelation(userID, targetUserID)
	if err != nil {
		return err
	}

	return s.relationRepo.Create(ctx, relation)
}

// Unfollow removes a relation between users
func (s *RelationServiceImpl) Unfollow(ctx context.Context, userID, targetUserID string) error {
	relation, err := domain.NewRelation(userID, targetUserID)
	if err != nil {
		return err
	}

	return s.relationRepo.Delete(ctx, relation)
}

// IsFollowing checks if a relation exists between users
func (s *RelationServiceImpl) IsFollowing(ctx context.Context, userID, targetUserID string) (*domain.RelationResponse, error) {
	relation, err := domain.NewRelation(userID, targetUserID)
	if err != nil {
		return nil, err
	}

	exists, err := s.relationRepo.Exists(ctx, relation)
	if err != nil {
		return &domain.RelationResponse{Status: false}, nil
	}

	return &domain.RelationResponse{Status: exists}, nil
}
