package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/jwt"
	"github.com/ccontreras/crispy-potato/models"
	"net/http"
	"time"
)

// Login handle the login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid user or password"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User email is a required field", 400)
		return
	}

	doc, exists := bd.LoginTry(t.Email, t.Password)
	if exists == false {
		http.Error(w, "Invalid user or password", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(doc)
	if err != nil {
		http.Error(w, "An error has occurred when we try to generate the token"+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    jwtKey,
		Expires:  expirationTime,
		Secure:   true,
		HttpOnly: false,
	})
}
