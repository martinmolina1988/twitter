package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*ConsultaBloqueo chequea si hay relacion entre 2 usuarios */
func ConsultaBloqueo(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("usuariobloqueadoid")
	UsuarioID := r.URL.Query().Get("usuarioid")

	var t models.Bloqueo
	t.UsuarioID = UsuarioID
	t.UsuarioBloqueadoID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoBloqueo(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
