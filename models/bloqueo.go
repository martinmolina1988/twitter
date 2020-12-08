package models

/*Bloqueo modelo para grabar la relacion de un usuario con otro */
type Bloqueo struct {
	UsuarioID          string `bson:"usuarioid" json:"usuarioId"`
	UsuarioBloqueadoID string `bson:"usuariobloqueadoid" json:"usuariobloqueadoId"`
}
