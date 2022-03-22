package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GeorgeHN666/proyect-Onlife-Server/db"
)

func SeeAllUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	typee := r.URL.Query().Get("type")
	pag := r.URL.Query().Get("pg")
	srch := r.URL.Query().Get("srch")

	if len(typee) < 1 {
		http.Error(w, "The param type should be sended", http.StatusBadRequest)
		return
	}

	if len(pag) < 1 {
		http.Error(w, "The param pg should be sended", http.StatusBadRequest)
		return
	}

	p, _ := strconv.Atoi(pag)

	pg := int64(p)

	results, good := db.SeeAllUsers(GlobalID, typee, srch, pg)
	if !good {
		http.Error(w, "Couldnt See All Users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(results)

}
