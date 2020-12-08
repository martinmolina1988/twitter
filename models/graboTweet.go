package models

import "time"

/*GraboTweet es el formato o estructura que tendrá nuestro Tweet en la BD */
type GraboTweet struct {
	UserID    string      `bson:"userid" json:"userid,omitempty"`
	Mensaje   string      `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha     time.Time   `bson:"fecha" json:"fecha,omitempty"`
	RTID      string      `bson:"rtid" json:"rtId,omitempty"`
	RTIDMsj   string      `bson:"rtidmsj" json:"rtIdmsj,omitempty"`
	RtaID     string      `bson:"rtaid" json:"rtaId,omitempty"`
	UserRTAID string      `bson:"userrtaid" json:"userrtaId,omitempty"`
	Likes     []Like      `bson:"like" json:"like,omitempty"`
	Retweet   []Retweet   `bson:"retweet" json:"retweet,omitempty"`
	Respuesta []Respuesta `bson:"respuesta" json:"respuesta,omitempty"`
}
