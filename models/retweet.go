package models

import "time"

/*Retweet es el formato o estructura que tendrá  el campo Tweet en la BD */
type Retweet struct {
	UserRTID string    `bson:"userrtid" json:"userrtid,omitempty"`
	Fecha    time.Time `bson:"fecha" json:"fecha,omitempty"`
}
