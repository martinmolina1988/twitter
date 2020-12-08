package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*LeoDenunciasPorCantidad lee los tweets de mis seguidores */
func LeoDenunciasPorCantidad(ID string, pagina int) ([]models.DenunciaCount, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	skip := (pagina - 1) * 20
	update := bson.M{
		"$group": bson.M{"_id": "$_id",
			"count": bson.M{"$sum": 1},
		},
	}

	condiciones := make([]bson.M, 0)
	condiciones = append(condiciones, bson.M{"$unwind": "$denuncias"})
	condiciones = append(condiciones, update)
	condiciones = append(condiciones, bson.M{"$sort": bson.M{"count": -1}})
	condiciones = append(condiciones, bson.M{"$skip": skip})
	condiciones = append(condiciones, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, condiciones)
	var result []models.DenunciaCount
	err = cursor.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
