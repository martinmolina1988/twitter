package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ConsultoLike consulta la relacion entre 2 usuarios */
func ConsultoLike(TweetID string, IDUsuario string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	objID, _ := primitive.ObjectIDFromHex(TweetID)

	condicion := bson.M{
		"_id":             objID,
		"like.userlikeid": IDUsuario,
	}
	var result models.Relacion
	err := col.FindOne(ctx, condicion).Decode(&result)

	if err != nil {
		return false, err
	}
	return true, nil
}
