package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Denuncias es la estructura con la que devolveremos las Denuncias */
type Denuncias struct {
	ID             primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID         string             `bson:"userid" json:"userId,omitempty"`
	UserDenunciaID string             `bson:"userdenunciaid" json:"userdenunciaId,omitempty"`
	TweetID        string             `bson:"tweetid" json:"tweetId,omitempty"`
	Fecha          time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
