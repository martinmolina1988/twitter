package models

import "time"

/*RTA es el formato o estructura que tendrá  el campo Tweet en la BD */
type RTA struct {
	UserRTAID string    `bson:"userrtaid" json:"userrtaId,omitempty"`
	UserID    string    `bson:"userid" json:"userId,omitempty"`
	Mensaje   string    `bson:"mensaje" json:"mensaje,omitempty"`
	TweetID   string    `bson:"tweetid" json:"tweetid,omitempty"`
	RTAID     string    `bson:"rtaid" json:"rtaid,omitempty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitempty"`
}
