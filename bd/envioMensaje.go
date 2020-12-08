package bd

import (
	"time"

	"context"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*EnvioMensaje envia mensajes al chat */
func EnvioMensaje(u models.Mensaje) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("mensajes")

	opts := options.Update().SetUpsert(true)

	query := bson.M{
		"userrecived": u.UserRecived,
		"usersend":    u.UserSend,
		"mensaje":     u.Mensaje,
	}

	update := bson.M{
		"$set": bson.M{"fecha": u.Fecha},
	}
	_, errr := col.UpdateOne(ctx, query, update, opts)
	if errr != nil {
		return false, errr
	}

	return true, nil
}
