package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DevuelvoTweetsSeguidores es la estructura con la que devolveremos los tweets */
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Tweet             struct {
		Mensaje   string      `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha     time.Time   `bson:"fecha" json:"fecha,omitempty"`
		ID        string      `bson:"_id" json:"_id,omitempty"`
		RTID      string      `bson:"rtid" json:"rtId,omitempty"`
		RTIDMsj   string      `bson:"rtidmsj" json:"rtIdmsj,omitempty"`
		RtaID     string      `bson:"rtaid" json:"rtaId,omitempty"`
		UserRTAID string      `bson:"userrtaid" json:"userrtaId,omitempty"`
		Likes     []Like      `bson:"like" json:"like,omitempty"`
		Retweet   []Retweet   `bson:"retweet" json:"retweet,omitempty"`
		Respuesta []Respuesta `bson:"respuesta" json:"respuesta,omitempty"`
	}
}
