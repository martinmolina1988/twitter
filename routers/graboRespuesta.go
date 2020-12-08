package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*GraboRespuesta permite grabar el tweet en la base de datos */
func GraboRespuesta(w http.ResponseWriter, r *http.Request) {
	var mensaje models.RTA
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	rtaid := mensaje.RTAID
	respuesta := models.Respuesta{
		UserRTAID: IDUsuario,
		Fecha:     time.Now(),
	}

	registro := models.GraboTweet{
		UserRTAID: mensaje.UserRTAID,
		RtaID:     mensaje.RTAID,
		UserID:    IDUsuario,
		Mensaje:   mensaje.Mensaje,
		Fecha:     time.Now(),
	}

	res := bd.InsertoRespuesta(registro, rtaid, respuesta)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

}
