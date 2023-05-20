package bd

import (
	"context"
	"time"

	"github.com/SteevenOrbe/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeoYaExisteUsuario recibe el email de parametro y chequea si ya esta en la BD.
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email:": email}

	var resultado models.Usuario

	//Valla a la colecion y haga la busquedad de un solo registro.
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	ID := resultado.ID.Hex() //convierte hexadecimal a string.

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID // va a devoolver toda la colecion del usuario.

}
