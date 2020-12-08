package models

import "time"

/*Mensaje es el formato o estructura que tendrá  el campo Tweet en la BD */
type Mensaje struct {
	UserSend    string    `bson:"usersend" json:"usersend,omitempty"`
	UserRecived string    `bson:"userrecived" json:"userrecived,omitempty"`
	Mensaje     string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha       time.Time `bson:"fecha" json:"fecha,omitempty"`
}
