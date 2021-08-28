package routers

import (
	"github.com/ccontreras/crispy-potato/bd"
	"io"
	"net/http"
	"os"
)

// GetAvatar for get the avatar
func GetAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "should send id parameter", http.StatusBadRequest)
		return
	}

	profile, err := bd.FindAProfile(ID)
	if err != nil {
		http.Error(w, "user not found", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)
	if err != nil {
		http.Error(w, "error copying the image", http.StatusBadRequest)
	}
}
