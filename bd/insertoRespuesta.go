package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoRespuesta graba el Tweet en la BD */
func InsertoRespuesta(rta models.GraboTweet, rtaID string, t models.Respuesta) (d models.IDD) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userrtaid": rta.UserRTAID,
		"rtaid":     rta.RtaID,
		"userid":    rta.UserID,
		"mensaje":   rta.Mensaje,
		"fecha":     rta.Fecha,
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
