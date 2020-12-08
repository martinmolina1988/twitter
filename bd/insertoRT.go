package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRT graba la relación en la BD */
func InsertoRT(rt models.GraboTweet, tweetID string, t models.Retweet) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")
	objID, _ := primitive.ObjectIDFromHex(tweetID)

	registro := bson.M{
		"userrtid": t.UserRTID,

		"fecha": t.Fecha,
	}

	update := bson.M{
		"$addToSet": bson.M{
			"retweet": registro,
		},
	}
	_, err := col.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return false, err
	}

	RT := bson.M{
		"userid": rt.UserID,
		"rtid":   rt.RTID,
		"fecha":  rt.Fecha,
	}
	_, err = col.InsertOne(ctx, RT)

	if err != nil {
		return false, err
	}

	return true, nil
}
