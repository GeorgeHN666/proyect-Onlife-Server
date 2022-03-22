package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func CreatePostEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var message models.Post

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Error trying to reade the json", http.StatusBadRequest)
		return
	}

	reg := models.Post{
		MessPost:   message.MessPost,
		DatePost:   time.Now(),
		UserIssued: GlobalID,
	}

	_, good, err := db.Post(reg)
	if err != nil || !good {
		http.Error(w, "Error creating the tweet", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
