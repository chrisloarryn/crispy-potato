package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"net/http"
)

// ReadRelation checks if there is or not relation between users
func ReadRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var response models.RelationResponse

	status, err := bd.ReadRelation(t)
	if err != nil || status == false {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
