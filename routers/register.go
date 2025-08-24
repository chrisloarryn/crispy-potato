package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
)

// Register is the function that creates any user in the database
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Received params error", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User email is a required field", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "User password should be at least 6 characters", 400)
		return
	}

	_, exists, _ := bd.CheckUserAlreadyExists(t.Email)
	if exists {
		http.Error(w, "The email you are trying to use is already taken by other one", 400)
		return
	}

	usr, status, err := bd.InsertRegister(t)
	if err != nil {
		http.Error(w, "An error has occurred when trying to insert the user to the database"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "User has not been inserted to the database", 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		http.Error(w, "An error has occurred when trying to encode the user to json"+err.Error(), 400)
		return
	}
}
