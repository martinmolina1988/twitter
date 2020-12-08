package models

/*RemuevoLike es el formato o estructura que tendrá nuestro Tweet en la BD */
type RemuevoLike struct {
	UserLikeID string `bson:"userlikeid" json:"userlikeid,omitempty"`
}
