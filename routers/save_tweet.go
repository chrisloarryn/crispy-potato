package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
)

// SaveTweet is for saving url handled
func SaveTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "An error has occurred when trying to decode"+err.Error(), 400)
		return
	}
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
	if !status {
		http.Error(w, "tweet has not been inserted", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
