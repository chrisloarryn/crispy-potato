package auth

import (
	"github.com/ccontreras/crispy-potato/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

// PasswordHasher implements the PasswordHasher interface using bcrypt
type PasswordHasher struct {
	cost int
}

// NewPasswordHasher creates a new PasswordHasher instance
func NewPasswordHasher() ports.PasswordHasher {
	return &PasswordHasher{
		cost: 8,
	}
}

// Hash hashes a password using bcrypt
func (h *PasswordHasher) Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(bytes), err
}

// Compare compares a hashed password with a plain text password
func (h *PasswordHasher) Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
