package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoTweett graba el Tweet en la BD */
func InsertoTweett(t models.GraboTweet) (d models.IDD) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	id := models.IDD{
		IDD: objID,
	}

	return id
}
