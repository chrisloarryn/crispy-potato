package bd

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the password for better security
func EncryptPassword(password string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
