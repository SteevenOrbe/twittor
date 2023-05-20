package main

import (
	"log"

	"github.com/SteevenOrbe/twittor/bd"
	"github.com/SteevenOrbe/twittor/handler"
)

// ESTO es donde esta nuestro ptoy en github ("github.com/SteevenOrbe/twittor) al momento de crear el repo.
func main() {
	//creheamos la conexion a mi bd
	if bd.ChequeoConnection() == 0 {
		log.Fatal("siN CONEXION A LA bd")
		return
	}
	handler.Manejadores()
}
