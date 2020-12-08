package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ConsultoRT consulta la relacion entre 2 usuarios */
func ConsultoRT(TweetID string, IDUsuario string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	objID, _ := primitive.ObjectIDFromHex(TweetID)

	condicion := bson.M{
		"_id":              objID,
		"retweet.userrtid": IDUsuario,
	}

	var resultado models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
