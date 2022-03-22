package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func UploadAvatarEndPoint(w http.ResponseWriter, r *http.Request) {

	imgFile, handler, _ := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var file string = "public/uploads/avatar/" + GlobalID + "." + extension

	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "There was an error uploading the image", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, imgFile)
	if err != nil {
		http.Error(w, "There was an error trying to copy the image", http.StatusBadRequest)
		return
	}

	var user models.User

	var good bool

	user.Avatar = GlobalID + "." + extension

	good, err = db.ModifyProfile(user, GlobalID)
	if err != nil || !good {
		http.Error(w, "Error Trying To Store The Image", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

}
