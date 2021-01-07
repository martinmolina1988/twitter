package bd

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func load() *options.ClientOptions {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	var ApplyURI = options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	return ApplyURI
}

/*MongoCN es el objeto de conexion a la BD*/
var MongoCN = ConectarBD()
var clientOption = load()

/*ConectarBD es la funcion que me perimite conectar la base de datos*/
func ConectarBD() *mongo.Client {
	load()
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Conexion exitosa a la base de datos")
	return client
}

/*ChequeoConnection es el ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
