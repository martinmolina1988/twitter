package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*AltaBloqueo realiza el registro de la relacion entre usuarios */
func AltaBloqueo(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Bloqueo
	t.UsuarioID = IDUsuario
	t.UsuarioBloqueadoID = ID

	status, err := bd.InsertoBloqueo(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
