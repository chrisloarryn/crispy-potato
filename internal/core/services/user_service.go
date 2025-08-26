package services

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// UserServiceImpl implements the UserService interface
type UserServiceImpl struct {
	userRepo    ports.UserRepository
	hasher      ports.PasswordHasher
	tokenGen    ports.TokenGenerator
	fileStorage ports.FileStorage
}

// NewUserService creates a new UserService instance
func NewUserService(
	userRepo ports.UserRepository,
	hasher ports.PasswordHasher,
	tokenGen ports.TokenGenerator,
	fileStorage ports.FileStorage,
) ports.UserService {
	return &UserServiceImpl{
		userRepo:    userRepo,
		hasher:      hasher,
		tokenGen:    tokenGen,
		fileStorage: fileStorage,
	}
}

// Register creates a new user
func (s *UserServiceImpl) Register(ctx context.Context, email, password string) (*domain.UserCreated, error) {
	// Check if user already exists
	_, exists, _ := s.userRepo.FindByEmail(ctx, email)
	if exists {
		return nil, fmt.Errorf("email is already taken")
	}

	// Create new user with domain validation
	user, err := domain.NewUser(email, password)
	if err != nil {
		return nil, err
	}

	// Hash password
	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	// Save user
	return s.userRepo.Create(ctx, user)
}

// Login authenticates a user
func (s *UserServiceImpl) Login(ctx context.Context, email, password string) (*domain.LoginResponse, error) {
	// Find user by email
	user, exists, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Compare password
	if err := s.hasher.Compare(user.Password, password); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate token
	token, err := s.tokenGen.Generate(user)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{Token: token}, nil
}

// GetProfile retrieves a user profile
func (s *UserServiceImpl) GetProfile(ctx context.Context, userID string) (*domain.User, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Remove sensitive data
	user.RemovePassword()
	return user, nil
}

// UpdateProfile updates user profile information
func (s *UserServiceImpl) UpdateProfile(ctx context.Context, userID string, name, surname, biographic, location, website string) error {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if err := user.Update(name, surname, biographic, location, website, user.Birthday); err != nil {
		return err
	}

	return s.userRepo.Update(ctx, user, userID)
}

// GetUsers retrieves a list of users
func (s *UserServiceImpl) GetUsers(ctx context.Context, currentUserID string, page int64, search, userType string) ([]*domain.User, error) {
	users, err := s.userRepo.FindAll(ctx, currentUserID, page, search, userType)
	if err != nil {
		return nil, err
	}

	// Remove sensitive data from all users
	for _, user := range users {
		user.RemovePassword()
	}

	return users, nil
}

// UploadAvatar uploads user avatar
func (s *UserServiceImpl) UploadAvatar(ctx context.Context, userID string, fileData []byte, fileName string) error {
	extension := strings.ToLower(filepath.Ext(fileName))
	avatarPath := fmt.Sprintf("uploads/avatars/%s%s", userID, extension)

	if err := s.fileStorage.Save(avatarPath, fileData); err != nil {
		return err
	}

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.SetAvatar(userID + extension)
	return s.userRepo.Update(ctx, user, userID)
}

// UploadBanner uploads user banner
func (s *UserServiceImpl) UploadBanner(ctx context.Context, userID string, fileData []byte, fileName string) error {
	extension := strings.ToLower(filepath.Ext(fileName))
	bannerPath := fmt.Sprintf("uploads/banners/%s%s", userID, extension)

	if err := s.fileStorage.Save(bannerPath, fileData); err != nil {
		return err
	}

	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.SetBanner(userID + extension)
	return s.userRepo.Update(ctx, user, userID)
}

// GetAvatar retrieves user avatar
func (s *UserServiceImpl) GetAvatar(ctx context.Context, userID string) ([]byte, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Avatar == "" {
		return nil, fmt.Errorf("avatar not found")
	}

	avatarPath := fmt.Sprintf("uploads/avatars/%s", user.Avatar)
	return s.fileStorage.Get(avatarPath)
}

// GetBanner retrieves user banner
func (s *UserServiceImpl) GetBanner(ctx context.Context, userID string) ([]byte, error) {
	user, err := s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user.Banner == "" {
		return nil, fmt.Errorf("banner not found")
	}

	bannerPath := fmt.Sprintf("uploads/banners/%s", user.Banner)
	return s.fileStorage.Get(bannerPath)
}
