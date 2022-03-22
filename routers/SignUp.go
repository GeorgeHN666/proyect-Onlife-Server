package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
	"github.com/GeorgeHN666/proyect-Onlife-Server/models"
)

func SignUpEndPoint(w http.ResponseWriter, r *http.Request) {

	var t models.User

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error Trying To Read The JSON Data", 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The Email Is Required", 400)
		return
	}

	if len(t.Pass) < 6 {
		http.Error(w, "The Password Must Contain at Least 6 Characters", 400)
		return
	}

	_, found, _ := db.UserExist(t.Email)
	if found {
		http.Error(w, "Theres Another Accout with This Email, Try to Login", 400)
		return
	}

	_, good, err := db.CreateUserInDB(t)

	if err != nil {
		http.Error(w, "There was An Error Creating The User", 400)
		return
	}

	if !good {
		http.Error(w, "Failed To Create User", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
