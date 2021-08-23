package bd

import (
	"github.com/ccontreras/crispy-potato/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginTry handle the login
func LoginTry(email string, password string) (models.User, bool) {
	user, found, _ := CheckUserAlreadyExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
