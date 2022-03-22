package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func PostComment(w http.ResponseWriter, r *http.Request) {

	postID := r.FormValue("id")
	message := r.FormValue("message")

	if len(postID) < 1 {
		http.Error(w, "The param id should be sended", http.StatusBadRequest)
		return
	}

	ps := models.Comments{
		PostID: postID,
		UserID: GlobalID,
		MessC:  message,
	}

	_, good, err := db.PostComment(ps)
	if err != nil || !good {
		http.Error(w, "There was an error trying to Insert the comment in db: line 27", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
