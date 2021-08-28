package routers

import (
	"github.com/ccontreras/crispy-potato/bd"
	"github.com/ccontreras/crispy-potato/models"
	"io"
	"net/http"
	"os"
	"strings"
)

// UploadAvatar to upload avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]
	var upFile string = "uploads/avatars/" + IDUser + "." + extension

	f, err := os.OpenFile(upFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "error uploading the image"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "error copying the image"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	user.Avatar = IDUser + "." + extension
	status, err = bd.ModifyRegister(user, IDUser)
	if err != nil || status == false {
		http.Error(w, "error saving the image"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
