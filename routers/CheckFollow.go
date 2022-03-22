package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func CheckFollowEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	userID := r.URL.Query().Get("id")

	if len(userID) < 1 {
		http.Error(w, "The params id should be sended", http.StatusBadRequest)
		return
	}

	var t models.Follow

	t.UserID = userID
	t.RelationID = GlobalID

	var results models.CheckRelation

	good, err := db.CheckFollow(t)
	if err != nil || !good {
		results.Status = false
	} else {
		results.Status = true
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(results)

}
