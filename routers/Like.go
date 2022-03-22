package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func LikeEndPoint(w http.ResponseWriter, r *http.Request) {

	var t models.Likes

	postID := r.URL.Query().Get("id")

	if len(postID) < 1 {
		http.Error(w, "The Param id Must be sended", http.StatusBadRequest)
		return
	}

	t.PostID = postID
	t.UserID = GlobalID

	_, good, err := db.CreateLike(t)
	if err != nil {
		http.Error(w, "Theres and error while trying to create a like", http.StatusBadRequest)
		return
	}

	if !good {
		http.Error(w, "Couldn't perform a like", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
