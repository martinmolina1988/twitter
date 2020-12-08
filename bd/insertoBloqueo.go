package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
)

/*InsertoBloqueo graba la relación en la BD */
func InsertoBloqueo(t models.Bloqueo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("bloqueos")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
