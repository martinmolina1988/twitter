package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*ModificarPassword modifica el password de usuario */
func ModificarPassword(w http.ResponseWriter, r *http.Request) {

	var t models.Password

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoPassword(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el password. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el password del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
