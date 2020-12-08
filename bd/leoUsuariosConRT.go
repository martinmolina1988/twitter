package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*LeoUsuariosConRT Lee los usuarios registrados en el sistema, si se recibe "R" en quienes
  trae solo los que se relacionan conmigo */
func LeoUsuariosConRT(ID string, page int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var results []*models.DevuelvoTweets

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}, {Key: "like.fecha", Value: -1}})
	opciones.SetSkip((page - 1) * 20)

	cur, err := col.Find(ctx, bson.M{"$or": []bson.M{{"userid": ID}, {"like.userlikeid": ID}}}, opciones)
	if err != nil {
		return results, false
	}

	for cur.Next(context.TODO()) {

		var registro models.DevuelvoTweets
		err := cur.Decode(&registro)
		if err != nil {
			return results, false
		}
		results = append(results, &registro)
	}
	return results, true

}
