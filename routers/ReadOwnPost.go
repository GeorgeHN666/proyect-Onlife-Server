package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func ReadOwnPostEndPoint(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "The Param id must be sended", http.StatusBadRequest)
		return
	}

	PG := r.URL.Query().Get("pg")

	if len(PG) < 1 {
		http.Error(w, "The Param pg Must Be Sended", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(PG)
	if err != nil {
		http.Error(w, "You Shoul Send The Param id As an integer", http.StatusBadRequest)
		return
	}

	pg := int64(page)

	response, good := db.ReadOwnPost(ID, pg)

	if !good {
		http.Error(w, "Fail To Read The Post", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}
