package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRTconMensaje graba la relación en la BD */
func InsertoRTconMensaje(rt models.GraboTweet) (d models.IDD) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	RT := bson.M{
		"userid":  rt.UserID,
		"mensaje": rt.Mensaje,
		"rtidmsj": rt.RTIDMsj,
		"fecha":   rt.Fecha,
	}
	result, err := col.InsertOne(ctx, RT)

	if err != nil {
		return
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	id := models.IDD{
		IDD: objID,
	}
	return id
}
