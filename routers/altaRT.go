package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*AltaRT realiza el registro de la relacion entre usuarios */
func AltaRT(w http.ResponseWriter, r *http.Request) {
	var mensaje models.GraboTweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if len(mensaje.RTID) < 1 {
		http.Error(w, "El parámetro rtid es obligatorio", http.StatusBadRequest)
		return
	}
	tweetid := mensaje.RTID
	retweet := models.Retweet{
		UserRTID: IDUsuario,
		Fecha:    time.Now(),
	}
	rt := models.GraboTweet{
		UserID: IDUsuario,
		RTID:   mensaje.RTID,
		Fecha:  time.Now(),
	}

	status, err := bd.InsertoRT(rt, tweetid, retweet)
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
