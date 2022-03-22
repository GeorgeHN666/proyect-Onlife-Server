package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func DeletePostEndPoint(w http.ResponseWriter, r *http.Request) {

	postID := r.URL.Query().Get("id")

	if len(postID) < 1 {
		http.Error(w, "The Param id Shoul Be Sended", http.StatusBadRequest)
		return
	}

	err := db.DeletePost(postID, GlobalID)
	if err != nil {
		http.Error(w, "There Was A problem Deleting The Post, Try Again Later", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
