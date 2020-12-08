package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoMensajes Lee los usuarios registrados en el sistema, si se recibe "R" en quienes
  trae solo los que se relacionan conmigo */
func LeoMensajes(usersend string, userrecived string, pagina int64) ([]*models.Mensaje, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("mensajes")

	var results []*models.Mensaje
	t := models.Mensaje{
		UserSend:    usersend,
		UserRecived: userrecived,
	}
	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: 1}, {Key: "fecha", Value: 1}})
	opciones.SetSkip((pagina - 1) * 20)

	condicion := bson.M{
		"$or": []bson.M{{
			"usersend":    t.UserSend,
			"userrecived": t.UserRecived,
		}, {
			"usersend":    t.UserRecived,
			"userrecived": t.UserSend,
		}},
	}

	cur, err := col.Find(ctx, condicion, opciones)
	if err != nil {
		return results, false
	}

	for cur.Next(context.TODO()) {

		var registro models.Mensaje
		err := cur.Decode(&registro)
		if err != nil {
			return results, false
		}
		results = append(results, &registro)
	}
	return results, true

}
