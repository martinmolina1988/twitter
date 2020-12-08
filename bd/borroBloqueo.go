package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
)

/*BorroBloqueo borra la relacion en la BD */
func BorroBloqueo(t models.Bloqueo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("bloqueos")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
