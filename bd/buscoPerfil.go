package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/SteevenOrbe/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// BuscoPerfil busca un perfil en la BD.
func BuscoPerfil(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)
	//Condicion de busquedad.
	condicion := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condicion).Decode(&perfil)

	//LA password no podemos mostrara al momento responder los datos de usuario.
	perfil.Password = ""
	if err != nil {
		fmt.Println("Registro no encontado " + err.Error())
		return perfil, err
	}
	// y si lo en contro nos respondera el perfil del usuario indicado eceto la password.
	return perfil, nil
}
