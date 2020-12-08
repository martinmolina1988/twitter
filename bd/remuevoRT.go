package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*RemuevoRT graba la relación en la BD */
func RemuevoRT(TweetID string, t models.RemuevoRetweet, UserID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(TweetID)

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	condicion := bson.M{
		"rtid":   TweetID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condicion)

	if err != nil {
		return false, err
	}

	registro := bson.M{
		"userrtid": t.UserRTID,
	}
	update := bson.M{
		"$pull": bson.M{
			"retweet": registro,
		}}

	_, err = col.UpdateOne(ctx, bson.M{"_id": objID}, update)

	if err != nil {
		return false, err
	}

	return true, nil
}
