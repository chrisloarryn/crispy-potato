package routers

import (
	"encoding/json"
	"net/http"

	"github.com/ccontreras/crispy-potato/bd"
)

// Me is a function to get the profile of current user
func Me(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Should send the id parameter", http.StatusBadRequest)
		return
	}
	profile, err := bd.FindAProfile(ID)
	if err != nil {
		http.Error(w, "An error has occurred when trying to find the profile"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
