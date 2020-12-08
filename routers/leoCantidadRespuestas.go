package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
)

/*LeoCantidadRespuestas Leo los tweets */
func LeoCantidadRespuestas(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.CantidadRespuestas(ID)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
