package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
)

/*AumentoNotif realiza el registro de la relacion entre usuarios */
func AumentoNotif(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parámetro TweetID es obligatorio", http.StatusBadRequest)
		return
	}

	status, err := bd.AumentoNotif(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el like "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el like "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
