package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func FollowEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("id")

	if len(userID) < 1 {
		http.Error(w, "The Param id must be sended", http.StatusBadRequest)
		return
	}

	var t models.Follow

	t.UserID = userID       // Who i will follow
	t.RelationID = GlobalID // who i am

	good, err := db.Follow(t)
	if err != nil || !good {
		http.Error(w, "There was an error trying to follow this user", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
