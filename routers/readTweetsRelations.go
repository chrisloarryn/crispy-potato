package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"net/http"
	"strconv"
)

// ReadTweetsRelations followers tweets
func ReadTweetsRelations(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "should send page parameter", http.StatusBadRequest)
		return
	}
	currPage, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "should send page parameter as integer greather than zero", http.StatusBadRequest)
		return
	}

	response, ok := bd.ReadTweetsFollowers(IDUser, currPage)
	if ok == false {
		http.Error(w, "error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
