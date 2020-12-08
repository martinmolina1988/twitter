package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Notificaciones es la estructura con la que devolveremos los Tweets */
type Notificaciones struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID      string             `bson:"userid" json:"userId,omitempty"`
	UserNotifID string             `bson:"usernotifid" json:"usernotifId,omitempty"`
	Tipo        string             `bson:"tipo" json:"tipo,omitempty"`
	TweetID     string             `bson:"tweetid" json:"tweetId,omitempty"`
	Fecha       time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
