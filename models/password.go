package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Password es el modelo del cambio de contraseña para el usuario de la base de MongoDB */
type Password struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`

	Password string `bson:"password" json:"password,omitempty"`
}
