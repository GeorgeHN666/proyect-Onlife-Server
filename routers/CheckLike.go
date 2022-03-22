package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func CheckLike(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	postID := r.URL.Query().Get("id")

	if len(postID) < 1 {
		http.Error(w, "The params id should be sended", http.StatusBadRequest)
		return
	}

	var t models.Likes

	t.PostID = postID
	t.UserID = GlobalID

	var response models.ResponseCheckLike

	good, err := db.CheckLike(t)
	if err != nil || !good {
		response.Status = false
	} else {
		response.Status = true
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}
