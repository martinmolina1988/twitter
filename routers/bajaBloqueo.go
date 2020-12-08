package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*BajaBloqueo realiza el borrado de la relacion entre usuarios */
func BajaBloqueo(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Bloqueo
	t.UsuarioID = IDUsuario
	t.UsuarioBloqueadoID = ID

	status, err := bd.BorroBloqueo(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
