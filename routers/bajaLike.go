package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*BajaLike realiza el registro de la relacion entre usuarios */
func BajaLike(w http.ResponseWriter, r *http.Request) {

	TweetID := r.URL.Query().Get("id")

	if len(TweetID) < 1 {
		http.Error(w, "El parámetro TweetID es obligatorio", http.StatusBadRequest)
		return
	}

	like := models.RemuevoLike{
		UserLikeID: IDUsuario,
	}

	status, err := bd.RemuevoLike(TweetID, like)
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
