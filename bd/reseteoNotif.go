package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*ReseteoNotif graba la relación en la BD */
func ReseteoNotif(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(ID)
	db := MongoCN.Database("twitter")
	col := db.Collection("usuarios")
	var num int64
	num = 0
	update := bson.M{
		"$set": bson.M{
			"notif": num,
		},
	}
	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
