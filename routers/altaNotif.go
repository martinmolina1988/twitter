package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*AltaNotif permite grabar el tweet en la base de datos */
func AltaNotif(w http.ResponseWriter, r *http.Request) {
	var t models.Notificaciones
	err := json.NewDecoder(r.Body).Decode(&t)

	registro := models.Notificaciones{
		UserNotifID: IDUsuario,
		Tipo:        t.Tipo,
		TweetID:     t.TweetID,
		UserID:      t.UserID,
		Fecha:       time.Now(),
	}

	status, err := bd.InsertoNotif(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro, reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
