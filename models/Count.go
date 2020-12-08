package models

/*Count captura del Body, el mensaje que nos llega */
type Count struct {
	Count int64 `bson:"count" json:"count"`
}
