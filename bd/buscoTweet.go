package bd

import (
	"context"
	"time"

	"github.com/martinmolina1988/twitter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BuscoTweet busca un perfil en la BD */
func BuscoTweet(ID string) (models.DevuelvoTweets, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("twitter")
	col := db.Collection("tweet")

	var tweet models.DevuelvoTweets
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&tweet)

	if err != nil {
		return tweet, err
	}
	return tweet, nil
}
