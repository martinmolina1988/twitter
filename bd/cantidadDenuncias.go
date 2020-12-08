package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*CantidadDenuncias lee los tweets de un perfil */
func CantidadDenuncias() (models.Count, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	var registro models.Count
	update := bson.M{
		"denuncias": bson.M{"$exists": "true"},
	}
	opts := options.Count()
	cursor, err := col.CountDocuments(ctx, update, opts)
	if err != nil {
		return registro, false
	}
	registro = models.Count{
		Count: cursor,
	}

	return registro, true
}
