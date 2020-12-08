package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CantidadRespuestas lee los tweets de un perfil */
func CantidadRespuestas(ID string) (models.Count, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	var registro models.Count
	condicion := bson.M{
		"rtaid": ID,
	}

	cursor, err := col.CountDocuments(ctx, condicion)
	if err != nil {
		return registro, false
	}
	registro = models.Count{
		Count: cursor,
	}

	return registro, true
}
