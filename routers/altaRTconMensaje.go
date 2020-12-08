package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*AltaRTconMensaje realiza el registro de la relacion entre usuarios */
func AltaRTconMensaje(w http.ResponseWriter, r *http.Request) {
	var mensaje models.GraboTweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if len(mensaje.RTIDMsj) < 1 {
		http.Error(w, "El parámetro rtidmsj es obligatorio", http.StatusBadRequest)
		return
	}

	rt := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		RTIDMsj: mensaje.RTIDMsj,
		Fecha:   time.Now(),
	}

	res := bd.InsertoRTconMensaje(rt)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}
