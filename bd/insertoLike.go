package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoLike graba la relación en la BD */
func InsertoLike(t models.Like, tweetID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(tweetID)
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userlikeid": t.UserLikeID,

		"fecha": t.Fecha,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"like": registro,
		},
	}
	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return false, err
	}

	return true, nil
}
