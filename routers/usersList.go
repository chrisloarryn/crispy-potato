package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"net/http"
	"strconv"
)

// UsersList capture params and send
func UsersList(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "should send page parameter greatger than zero", http.StatusBadRequest)
		return
	}
	int64Page := int64(pageTemp)

	result, ok := bd.ReadAllUsers(IDUser, int64Page, search, typeUser)
	if ok == false || !ok {
		http.Error(w, "Error reading users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
