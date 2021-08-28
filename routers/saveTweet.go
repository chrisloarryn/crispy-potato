package routers

import (
	"encoding/json"
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"net/http"
	"time"
)

// SaveTweet is for saving url handled
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	register := models.SaveTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}
	_, status, err := bd.InsertTweet(register)

	if err != nil {
		http.Error(w, "An error has occurred when trying to insert to mongodb"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "tweet has not been inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
