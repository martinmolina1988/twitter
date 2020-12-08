package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/martinmolina1988/twitter/bd"
)

/*LeoTweetsTotal Leo los tweets */
func LeoTweetsTotal(w http.ResponseWriter, r *http.Request) {

	Mensaje := r.URL.Query().Get("mensaje")

	if len(Mensaje) < 1 {
		http.Error(w, "Debe enviar el parámetro mensaje", http.StatusBadRequest)
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
	respuesta, correcto := bd.LeoTweetsTotal(Mensaje, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)
}
