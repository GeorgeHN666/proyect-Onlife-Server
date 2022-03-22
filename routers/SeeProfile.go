package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func SeeProfileEndPoint(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if len(r.URL.Query().Get("id")) == 0 {
		http.Error(w, "You Should Send The Param id", 400)
		return
	}

	profile, err := db.SeeProfile(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "The user doesn't exist", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(profile)

}
