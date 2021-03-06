package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/martinmolina1988/twitter/bd"
)

/*LeoUltimoMensaje Leo los tweets */
func LeoUltimoMensaje(w http.ResponseWriter, r *http.Request) {

	usersend := r.URL.Query().Get("usersend")
	if len(usersend) < 1 {
		http.Error(w, "Debe enviar el parámetro usersend", http.StatusBadRequest)
		return
	}
	userrecived := r.URL.Query().Get("userrecived")
	if len(userrecived) < 1 {
		http.Error(w, "Debe enviar el parámetro userrecived", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro página", http.StatusBadRequest)
		return
	}
	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoUltimoMensaje(usersend, userrecived, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
