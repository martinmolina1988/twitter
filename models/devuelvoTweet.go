package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweets es la estructura con la que devolveremos los Tweets */
type DevuelvoTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID    string             `bson:"userid" json:"userId,omitempty"`
	Mensaje   string             `bson:"mensaje" json:"mensaje,omitempty"`
	RTID      string             `bson:"rtid" json:"rtId,omitempty"`
	RTIDMsj   string             `bson:"rtidmsj" json:"rtIdmsj,omitempty"`
	RtaID     string             `bson:"rtaid" json:"rtaId,omitempty"`
	UserRTAID string             `bson:"userrtaid" json:"userrtaId,omitempty"`
	Fecha     time.Time          `bson:"fecha" json:"fecha,omitempty"`
	Likes     []Like             `bson:"like" json:"like,omitempty"`
	Retweet   []Retweet          `bson:"retweet" json:"retweet,omitempty"`
	Respuesta []Respuesta        `bson:"respuesta" json:"respuesta,omitempty"`
}
