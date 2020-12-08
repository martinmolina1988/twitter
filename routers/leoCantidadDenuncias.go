package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
)

/*LeoCantidadDenuncias Leo los tweets */
func LeoCantidadDenuncias(w http.ResponseWriter, r *http.Request) {

	respuesta, correcto := bd.CantidadDenuncias()
	if correcto == false {
		http.Error(w, "Error al leer las denuncias", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
