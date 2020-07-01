package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Usuario es el modelo de la base de mongodb*/
type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempy" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempy"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempy"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempy"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempy"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempy"`
	Banner          string             `bson:"banner" json:"banner,omitempy"`
	Biografia       string             `bson:"biogriafia" json:"biogriafia,omitempy"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempy"`
	SitioWeb        string             `bson:"sitioWeb" json:"sitioWeb,omitempy"`
}
