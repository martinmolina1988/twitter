package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*AumentoNotif graba la relación en la BD */
func AumentoNotif(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(ID)
	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")

	update := bson.M{
		"$inc": bson.M{
			"notif": 1,
		},
	}
	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
