package routers

import (
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
)

/*EliminarNotif permite borrar un Tweet determinado */
func EliminarNotif(w http.ResponseWriter, r *http.Request) {
	UserID := r.URL.Query().Get("usernotifid")
	Tipo := r.URL.Query().Get("tipo")
	TweetID := r.URL.Query().Get("tweetid")

	err := bd.BorroNotif(TweetID, UserID, Tipo)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
