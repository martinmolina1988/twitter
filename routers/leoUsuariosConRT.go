package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/martinmolina1988/twitter/bd"
)

/*LeoUsuariosConRTs leo la lista de los usuarios */
func LeoUsuariosConRTs(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	page := r.URL.Query().Get("page")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := bd.LeoUsuariosConRT(ID, pag)
	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
