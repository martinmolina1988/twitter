package models

/*RemuevoRespuesta es el formato o estructura que tendrá nuestro Tweet en la BD */
type RemuevoRespuesta struct {
	UserRTAID string `bson:"userrtaid" json:"userrtaId,omitempty"`
}
