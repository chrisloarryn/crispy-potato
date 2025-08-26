package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserRepository implements UserRepository for testing
type MockUserRepository struct {
	users         map[string]*domain.User
	shouldFail    bool
	existingEmail string
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*domain.User),
	}
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) (*domain.UserCreated, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}

	user.ID = primitive.NewObjectID()
	m.users[user.Email] = user

	return &domain.UserCreated{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*domain.User, bool, error) {
	if m.shouldFail {
		return nil, false, errors.New("database error")
	}

	user, exists := m.users[email]
	if !exists && email == m.existingEmail {
		// Simulate existing user
		user = &domain.User{
			ID:    primitive.NewObjectID(),
			Email: email,
		}
		return user, true, nil
	}

	return user, exists, nil
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*domain.User, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}

	for _, user := range m.users {
		if user.ID.Hex() == id {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User, id string) error {
	if m.shouldFail {
		return errors.New("database error")
	}

	for email, existingUser := range m.users {
		if existingUser.ID.Hex() == id {
			m.users[email] = user
			return nil
		}
	}

	return errors.New("user not found")
}

func (m *MockUserRepository) FindAll(ctx context.Context, currentUserID string, page int64, search, userType string) ([]*domain.User, error) {
	if m.shouldFail {
		return nil, errors.New("database error")
	}

	var result []*domain.User
	for _, user := range m.users {
		if user.ID.Hex() != currentUserID {
			result = append(result, user)
		}
	}

	return result, nil
}

// MockPasswordHasher implements PasswordHasher for testing
type MockPasswordHasher struct {
	shouldFail bool
}

func NewMockPasswordHasher() *MockPasswordHasher {
	return &MockPasswordHasher{}
}

func (m *MockPasswordHasher) Hash(password string) (string, error) {
	if m.shouldFail {
		return "", errors.New("hash error")
	}
	return "hashed_" + password, nil
}

func (m *MockPasswordHasher) Compare(hashedPassword, password string) error {
	if m.shouldFail {
		return errors.New("compare error")
	}

	expected := "hashed_" + password
	if hashedPassword != expected {
		return errors.New("invalid password")
	}

	return nil
}

// MockTokenGenerator implements TokenGenerator for testing
type MockTokenGenerator struct {
	shouldFail bool
}

func NewMockTokenGenerator() *MockTokenGenerator {
	return &MockTokenGenerator{}
}

func (m *MockTokenGenerator) Generate(user *domain.User) (string, error) {
	if m.shouldFail {
		return "", errors.New("token generation error")
	}
	return "mock_token_" + user.Email, nil
}

func (m *MockTokenGenerator) Validate(token string) (*domain.User, error) {
	if m.shouldFail {
		return nil, errors.New("invalid token")
	}

	return &domain.User{Email: "test@test.com"}, nil
}

// MockFileStorage implements FileStorage for testing
type MockFileStorage struct {
	files      map[string][]byte
	shouldFail bool
}

func NewMockFileStorage() *MockFileStorage {
	return &MockFileStorage{
		files: make(map[string][]byte),
	}
}

func (m *MockFileStorage) Save(path string, data []byte) error {
	if m.shouldFail {
		return errors.New("storage error")
	}
	m.files[path] = data
	return nil
}

func (m *MockFileStorage) Get(path string) ([]byte, error) {
	if m.shouldFail {
		return nil, errors.New("storage error")
	}

	data, exists := m.files[path]
	if !exists {
		return nil, errors.New("file not found")
	}

	return data, nil
}

func (m *MockFileStorage) Delete(path string) error {
	if m.shouldFail {
		return errors.New("storage error")
	}
	delete(m.files, path)
	return nil
}

// Test cases
func TestUserService_Register_Success(t *testing.T) {
	// Arrange
	userRepo := NewMockUserRepository()
	hasher := NewMockPasswordHasher()
	tokenGen := NewMockTokenGenerator()
	storage := NewMockFileStorage()

	service := services.NewUserService(userRepo, hasher, tokenGen, storage)

	// Act
	result, err := service.Register(context.Background(), "test@test.com", "password123")

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected result, got nil")
	}

	if result.Email != "test@test.com" {
		t.Errorf("Expected email 'test@test.com', got %s", result.Email)
	}
}

func TestUserService_Register_UserAlreadyExists(t *testing.T) {
	// Arrange
	userRepo := NewMockUserRepository()
	userRepo.existingEmail = "existing@test.com"
	hasher := NewMockPasswordHasher()
	tokenGen := NewMockTokenGenerator()
	storage := NewMockFileStorage()

	service := services.NewUserService(userRepo, hasher, tokenGen, storage)

	// Act
	result, err := service.Register(context.Background(), "existing@test.com", "password123")

	// Assert
	if err == nil {
		t.Error("Expected error for existing user")
	}

	if result != nil {
		t.Error("Expected nil result for existing user")
	}
}

func TestUserService_Register_InvalidEmail(t *testing.T) {
	// Arrange
	userRepo := NewMockUserRepository()
	hasher := NewMockPasswordHasher()
	tokenGen := NewMockTokenGenerator()
	storage := NewMockFileStorage()

	service := services.NewUserService(userRepo, hasher, tokenGen, storage)

	// Act
	result, err := service.Register(context.Background(), "", "password123")

	// Assert
	if err == nil {
		t.Error("Expected error for empty email")
	}

	if result != nil {
		t.Error("Expected nil result for invalid email")
	}
}

func TestUserService_Register_WeakPassword(t *testing.T) {
	// Arrange
	userRepo := NewMockUserRepository()
	hasher := NewMockPasswordHasher()
	tokenGen := NewMockTokenGenerator()
	storage := NewMockFileStorage()

	service := services.NewUserService(userRepo, hasher, tokenGen, storage)

	// Act
	result, err := service.Register(context.Background(), "test@test.com", "123")

	// Assert
	if err == nil {
		t.Error("Expected error for weak password")
	}

	if result != nil {
		t.Error("Expected nil result for weak password")
	}
}

func TestUserService_Login_Success(t *testing.T) {
	// Arrange
	userRepo := NewMockUserRepository()
	hasher := NewMockPasswordHasher()
	tokenGen := NewMockTokenGenerator()
	storage := NewMockFileStorage()

	// Create a user first
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@test.com",
		Password: "hashed_password123",
	}
	userRepo.users["test@test.com"] = user

	service := services.NewUserService(userRepo, hasher, tokenGen, storage)

	// Act
	result, err := service.Login(context.Background(), "test@test.com", "password123")

	// Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Error("Expected result, got nil")
	}

	if result.Token == "" {
		t.Error("Expected token, got empty string")
	}
}
