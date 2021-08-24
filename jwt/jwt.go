package jwt

import (
	"github.com/ccontreras/crispy-potato/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// GenerateJWT is the function to generate JWT
func GenerateJWT(t models.User) (string, error) {
	myKey := []byte("MastersOfDevelopment_facebookGroup")

	payload := jwt.MapClaims{
		"email": t.Email,
		"name": t.Name,
		"surname": t.Surname,
		"birthday": t.Birthday,
		"biographic": t.Biographic,
		"location": t.Location,
		"website": t.Website,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}