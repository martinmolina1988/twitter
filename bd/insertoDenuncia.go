package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoDenuncia graba la Denuncia en la BD */
func InsertoDenuncia(t models.Denuncias) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("denuncias")

	registro := bson.M{
		"userdenunciaid": t.UserDenunciaID,
		"fecha":          t.Fecha,
		"tweetid":        t.TweetID,
		"userid":         t.UserID,
	}

	_, err := col.InsertOne(ctx, registro)
	if err != nil {
		return false, err
	}

	return true, nil
}
