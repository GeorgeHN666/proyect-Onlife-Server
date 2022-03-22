package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func ModifyUserEndPoint(w http.ResponseWriter, r *http.Request) {

	var t models.User

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Theres An Error Trying To Read The Data", 400)
		return
	}

	var results bool

	results, err = db.ModifyProfile(t, GlobalID)
	if err != nil {
		http.Error(w, "There was an error Trying to modify The Profile", 400)
		return
	}

	if !results {
		http.Error(w, "Couldn't Modify The Profile", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
