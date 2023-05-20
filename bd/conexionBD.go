package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es el obj de conecion
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:developer@twittor.mlzspof.mongodb.net/?retryWrites=true&w=majority")

// ConnectarBD es la funcion que me permite conectar a  la BD
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil) //para ver si esta disponible la bd.
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion exitosa con la BD..")
	return client
}

// ChequeoConnection  es le ping a la BD
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil) //para ver si esta disponible la bd.
	if err != nil {
		return 0 // si hay error
	}
	return 1 //no hay error
}
