package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func ObtainBannerEndpoint(w http.ResponseWriter, r *http.Request) {

	userID := r.URL.Query().Get("id")

	if len(userID) < 1 {
		http.Error(w, "The params id should be sended", http.StatusBadRequest)
		return
	}

	profile, err := db.SeeProfile(userID)
	if err != nil {
		http.Error(w, "The user wasn't found", http.StatusBadRequest)
		return
	}

	img, err := os.Open("public/uploads/banner/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Image Not Found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, img)
	if err != nil {
		http.Error(w, "Error Copying The Image", http.StatusBadRequest)
		return
	}

}
