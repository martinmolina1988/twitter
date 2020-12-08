package models

/*RemuevoRetweet es el formato o estructura que tendrá nuestro Tweet en la BD */
type RemuevoRetweet struct {
	UserRTID string `bson:"userrtid" json:"userrtid,omitempty"`
}
