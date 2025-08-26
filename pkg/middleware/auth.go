package middleware

import (
	"context"
	"net/http"

	"github.com/ccontreras/crispy-potato/internal/core/ports"
)

// AuthMiddleware handles JWT authentication
type AuthMiddleware struct {
	tokenGenerator ports.TokenGenerator
	userRepo       ports.UserRepository
}

// NewAuthMiddleware creates a new AuthMiddleware instance
func NewAuthMiddleware(tokenGenerator ports.TokenGenerator, userRepo ports.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		tokenGenerator: tokenGenerator,
		userRepo:       userRepo,
	}
}

// ValidateJWT validates JWT tokens and adds user context
func (m *AuthMiddleware) ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		user, err := m.tokenGenerator.Validate(token)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		// Get user from database to ensure they still exist and get current info
		userFromDB, exists, err := m.userRepo.FindByEmail(r.Context(), user.Email)
		if err != nil || !exists {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		// Add user ID to request context
		ctx := context.WithValue(r.Context(), "userID", userFromDB.ID.Hex())
		ctx = context.WithValue(ctx, "email", userFromDB.Email)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}
