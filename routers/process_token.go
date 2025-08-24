package routers

import (
	"errors"
	"strings"

	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"github.com/dgrijalva/jwt-go"
)

// IDUser user id used in all the endpoints
var IDUser string

// Email user email used in all the endpoints
var Email string

// ProcessToken function to process the token
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("MastersOfDevelopment_facebookGroup")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token format invalid")
	}
	token = strings.TrimSpace(splitToken[1])
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		userFound, found, _ := bd.CheckUserAlreadyExists(claims.Email)
		if found {
			Email = claims.Email
			IDUser = userFound.ID.Hex()
		}
		return claims, found, IDUser, nil
	}

	if !jwtToken.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
