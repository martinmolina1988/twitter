package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoDen graba la relación en la BD */
func InsertoDen(t models.Denuncias) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tweetid := t.TweetID
	objID, _ := primitive.ObjectIDFromHex(tweetid)
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userdenunciaid": t.UserDenunciaID,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"denuncias": registro,
		},
	}
	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return false, err
	}

	updateDate := bson.M{
		"$set": bson.M{"fechaDenuncia": t.Fecha},
	}
	_, errr := col.UpdateOne(ctx, bson.M{"_id": objID}, updateDate)
	if errr != nil {
		return false, errr
	}

	return true, nil
}
