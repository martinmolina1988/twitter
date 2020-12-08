package bd

import (
	"time"

	"context"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MensajeUsuario permite modificar el perfil del usuario */
func MensajeUsuario(u models.Mensaje, v models.Mensaje) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("usuariosMensajes")

	opts := options.Update().SetUpsert(true)

	query := bson.M{
		"userrecived": u.UserRecived,
		"usersend":    u.UserSend,
	}

	update := bson.M{
		"$set": bson.M{"fecha": u.Fecha},
	}
	_, errr := col.UpdateOne(ctx, query, update, opts)
	if errr != nil {
		return false, errr
	}
	querys := bson.M{
		"userrecived": v.UserRecived,
		"usersend":    v.UserSend,
	}

	updates := bson.M{
		"$set": bson.M{"fecha": v.Fecha},
	}
	_, err := col.UpdateOne(ctx, querys, updates, opts)
	if err != nil {
		return false, err
	}

	return true, nil
}
