package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"net/http"
	"strconv"
)

// ReadTweets for reading tweets
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "should send id parameter", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "should send page parameter", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "should send page parameter with number upper than zero", http.StatusBadRequest)
		return
	}

	int64Page := int64(page)

	response, ok := bd.ReadTweets(ID, int64Page)
	if ok == false {
		http.Error(w, "error reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)
}
