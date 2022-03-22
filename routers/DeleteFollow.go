package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func DeleteFollowEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("id")

	if len(userID) < 1 {
		http.Error(w, "The param id should be sended", http.StatusBadRequest)
		return
	}

	var t models.Follow

	t.UserID = userID
	t.RelationID = GlobalID

	good, err := db.DeleteFollow(t)
	if err != nil || !good {
		http.Error(w, "Fail to delete the relation", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
