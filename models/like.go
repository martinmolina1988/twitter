package models

import "time"

/*Like es el formato o estructura que tendrá  el campo like en la BD */
type Like struct {
	UserLikeID string    `bson:"userlikeid" json:"userlikeid,omitempty"`
	Fecha      time.Time `bson:"fecha" json:"fecha,omitempty"`
}
