package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
)

/*ReseteoNotif realiza el registro de la relacion entre usuarios */
func ReseteoNotif(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	status, err := bd.ReseteoNotif(ID)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar resetear la notificacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado resetear la notificacion "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
