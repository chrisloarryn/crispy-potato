package auth

import (
	"errors"
	"strings"
	"time"

	"github.com/ccontreras/crispy-potato/internal/core/domain"
	"github.com/ccontreras/crispy-potato/internal/core/ports"
	"github.com/dgrijalva/jwt-go"
)

// JWTTokenGenerator implements the TokenGenerator interface
type JWTTokenGenerator struct {
	secretKey []byte
}

// Claims represents JWT claims
type Claims struct {
	Email string `json:"email"`
	ID    string `json:"_id"`
	jwt.StandardClaims
}

// NewJWTTokenGenerator creates a new JWTTokenGenerator instance
func NewJWTTokenGenerator(secretKey string) ports.TokenGenerator {
	return &JWTTokenGenerator{
		secretKey: []byte(secretKey),
	}
}

// Generate generates a JWT token for a user
func (g *JWTTokenGenerator) Generate(user *domain.User) (string, error) {
	payload := jwt.MapClaims{
		"email":      user.Email,
		"name":       user.Name,
		"surname":    user.Surname,
		"birthday":   user.Birthday,
		"biographic": user.Biographic,
		"location":   user.Location,
		"website":    user.Website,
		"_id":        user.ID.Hex(),
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return token.SignedString(g.secretKey)
}

// Validate validates a JWT token and returns the user information
func (g *JWTTokenGenerator) Validate(tokenString string) (*domain.User, error) {
	claims := &Claims{}

	splitToken := strings.Split(tokenString, "Bearer")
	if len(splitToken) != 2 {
		return nil, errors.New("token format invalid")
	}
	tokenString = strings.TrimSpace(splitToken[1])

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return g.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	// Create a user object from claims
	user := &domain.User{
		Email: claims.Email,
	}

	return user, nil
}
