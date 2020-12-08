package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoBloqueo consulta la relacion entre 2 usuarios */
func ConsultoBloqueo(t models.Bloqueo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("bloqueos")

	condicion := bson.M{
		"usuarioid":          t.UsuarioID,
		"usuariobloqueadoid": t.UsuarioBloqueadoID,
	}

	var resultado models.Bloqueo
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
