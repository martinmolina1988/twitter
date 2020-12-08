package models

import "time"

/*Respuesta es el formato o estructura que tendrá el campo respuesta en la BD */
type Respuesta struct {
	UserRTAID string    `bson:"userrtaid" json:"userrtaid,omitempty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitempty"`
}
