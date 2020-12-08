package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*ConsultaRT chequea si hay relacion entre 2 usuarios */
func ConsultaRT(w http.ResponseWriter, r *http.Request) {
	TweetID := r.URL.Query().Get("id")

	var resp models.RespuestaConsultaRelacion
	status, err := bd.ConsultoRT(TweetID, IDUsuario)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
