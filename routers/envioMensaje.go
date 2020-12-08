package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*EnvioMensaje modifica el perfil de usuario */
func EnvioMensaje(w http.ResponseWriter, r *http.Request) {

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
		Mensaje:     t.Mensaje,
		UserSend:    IDUsuario,
	}

	var status bool

	status, err = bd.EnvioMensaje(registro)
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
