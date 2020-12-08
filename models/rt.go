package models

import "time"

/*RT es el formato o estructura que tendrá  el campo Tweet en la BD */
type RT struct {
	UserID  string    `bson:"userid" json:"userId,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	TweetID string    `bson:"tweetid" json:"tweetid,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
