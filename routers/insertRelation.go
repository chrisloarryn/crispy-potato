package routers

import (
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"net/http"
)

// InsertRelation do the registry
func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "should send id parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "an error has occurred when inserting relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "relation has not been inserted "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
