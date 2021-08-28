package routers

import (
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"net/http"
)

// DeleteRelation for delete
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "should send id parameter", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	status, err := bd.DeleteRelation(t)
	if err != nil {
		http.Error(w, "an error has occurred when deleting relation "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "relation has not been deleted "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
