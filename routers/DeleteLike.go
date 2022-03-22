package routers

import (
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func DeleteLikeEndpoint(w http.ResponseWriter, r *http.Request) {

	likeID := r.URL.Query().Get("id")

	if len(likeID) < 1 {
		http.Error(w, "The Param id Should Be Sended", http.StatusBadRequest)
		return
	}

	good, err := db.DeleteLike(likeID, GlobalID)
	if err != nil || !good {
		http.Error(w, "There was an error trying to delete the like", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
