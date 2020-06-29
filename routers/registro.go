package routers

import (
	"encoding/json"
	"net/http"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/models"
)

/*Registro es la funcion para crear en la BD el registro de ususario*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El mail de usuario es requerido", 400)
	}

	if len(t.Password) < 0 {
		http.Error(w, "Debe indiciar una contraseña de al menos 6 caracteres", 400)
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario con ese mail", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usaurio", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
