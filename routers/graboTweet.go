package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*GraboTweet permite grabar el tweet en la base de datos */
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	respuesta := bd.InsertoTweett(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
