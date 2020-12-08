package models

/*Email es el modelo del cambio de email para el usuario de la base de MongoDB */
type Email struct {
	Email string `bson:"email" json:"email,omitempty"`
}
