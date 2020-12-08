package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*BorroNotif borra un tweet determinado */
func BorroNotif(TweetID string, UserID string, Tipo string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("notificaciones")

	condicion := bson.M{
		"usernotifid": UserID,
		"tipo":        Tipo,
		"tweetid":     TweetID,
	}

	_, err := col.DeleteOne(ctx, condicion)
	return err
}
