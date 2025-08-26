package domain

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user entity in the domain
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name       string             `bson:"name" json:"name,omitempty"`
	Surname    string             `bson:"surname" json:"surname,omitempty"`
	Birthday   time.Time          `bson:"birthday" json:"birthday,omitempty"`
	Email      string             `bson:"email" json:"email"`
	Password   string             `bson:"password" json:"password,omitempty"`
	Avatar     string             `bson:"avatar" json:"avatar,omitempty"`
	Banner     string             `bson:"banner" json:"banner,omitempty"`
	Biographic string             `bson:"biographic" json:"biographic,omitempty"`
	Location   string             `bson:"location" json:"location,omitempty"`
	Website    string             `bson:"website" json:"website,omitempty"`
}

// UserCreated represents the response when a user is created
type UserCreated struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email string             `bson:"email" json:"email"`
}

// NewUser creates a new user with validation
func NewUser(email, password string) (*User, error) {
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	return &User{
		Email:    email,
		Password: password,
	}, nil
}

// Update updates user fields with validation
func (u *User) Update(name, surname, biographic, location, website string, birthday time.Time) error {
	u.Name = name
	u.Surname = surname
	u.Biographic = biographic
	u.Location = location
	u.Website = website
	u.Birthday = birthday
	return nil
}

// SetAvatar sets the user avatar
func (u *User) SetAvatar(avatar string) {
	u.Avatar = avatar
}

// SetBanner sets the user banner
func (u *User) SetBanner(banner string) {
	u.Banner = banner
}

// RemovePassword removes password from user (for safe responses)
func (u *User) RemovePassword() {
	u.Password = ""
}

// validateEmail validates email format
func validateEmail(email string) error {
	if len(email) == 0 {
		return errors.New("email is required")
	}
	// Basic email validation - could be enhanced with regex
	return nil
}

// validatePassword validates password strength
func validatePassword(password string) error {
	if len(password) < 6 {
		return errors.New("password should be at least 6 characters")
	}
	return nil
}
