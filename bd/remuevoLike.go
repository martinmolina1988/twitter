package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RemuevoLike graba la relación en la BD */
func RemuevoLike(TweetID string, t models.RemuevoLike) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(TweetID)

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	registro := bson.M{
		"userlikeid": t.UserLikeID,
	}

	update := bson.M{
		"$pull": bson.M{
			"like": registro,
		}}

	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil {
		return false, err
	}

	return true, nil
}
