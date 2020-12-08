package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*IDD captura del Body, el mensaje que nos llega */
type IDD struct {
	IDD primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
}
