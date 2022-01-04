package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "incorrect parameters"+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModifyRegister(t, IDUser)
	if err != nil {
		http.Error(w, "an error has occurred when trying to update, please try again"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "user has not been modified yet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
