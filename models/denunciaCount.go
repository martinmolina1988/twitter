package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*DenunciaCount captura del Body, el mensaje que nos llega */
type DenunciaCount struct {
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	Count int                `bson:"count" json:"count"`
}
