package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*AltaDenuncia permite grabar la Denuncia en la base de datos */
func AltaDenuncia(w http.ResponseWriter, r *http.Request) {
	var t models.Denuncias
	err := json.NewDecoder(r.Body).Decode(&t)

	registro := models.Denuncias{
		UserDenunciaID: IDUsuario,
		TweetID:        t.TweetID,
		UserID:         t.UserID,
		Fecha:          time.Now(),
	}

	status, err := bd.InsertoDen(registro)
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
