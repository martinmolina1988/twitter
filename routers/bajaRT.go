package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*BajaRT realiza el registro de la relacion entre usuarios */
func BajaRT(w http.ResponseWriter, r *http.Request) {

	TweetID := r.URL.Query().Get("id")

	if len(TweetID) < 1 {
		http.Error(w, "El parámetro TweetID es obligatorio", http.StatusBadRequest)
		return
	}
	retweet := models.RemuevoRetweet{
		UserRTID: IDUsuario,
	}
	status, err := bd.RemuevoRT(TweetID, retweet, IDUsuario)
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
