package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*MensajeUsuario modifica el perfil de usuario */
func MensajeUsuario(w http.ResponseWriter, r *http.Request) {

	var t models.Mensaje

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}
	fecha := time.Now()

	registro := models.Mensaje{
		Fecha:       fecha,
		UserRecived: t.UserRecived,
		UserSend:    IDUsuario,
	}

	registros := models.Mensaje{
		Fecha:       fecha,
		UserRecived: IDUsuario,
		UserSend:    t.UserRecived,
	}

	var status bool

	status, err = bd.MensajeUsuario(registro, registros)
	if err != nil {
		http.Error(w, "Ocurrión un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado modificar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
