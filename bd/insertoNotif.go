package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*InsertoNotif graba la relación en la BD */
func InsertoNotif(t models.Notificaciones) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("notificaciones")

	registro := bson.M{
		"usernotifid": t.UserNotifID,
		"tipo":        t.Tipo,
		"fecha":       t.Fecha,
		"tweetid":     t.TweetID,
		"userid":      t.UserID,
	}

	_, err := col.InsertOne(ctx, registro)
	if err != nil {
		return false, err
	}

	return true, nil
}
