package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	commentID := r.URL.Query().Get("id")

	if len(commentID) < 1 {
		http.Error(w, "The param id should be sended", http.StatusBadRequest)
		return
	}

	good, err := db.DeleteComment(commentID, GlobalID)
	if err != nil || !good {
		http.Error(w, "Couldn't  Delete The Comment", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
