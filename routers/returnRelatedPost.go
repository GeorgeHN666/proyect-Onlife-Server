package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func SeeRelatedPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	pag := r.URL.Query().Get("pg")

	if len(pag) < 1 {
		http.Error(w, "The param pg should be sended", http.StatusBadRequest)
		return
	}

	p, _ := strconv.Atoi(pag)

	pg := int64(p)

	response, good := db.ReadAllPost(GlobalID, pg)
	if !good {
		http.Error(w, "There something Wrong With Database", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(response)

}
