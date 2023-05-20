package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/SteevenOrbe/twittor/middlew"
	"github.com/SteevenOrbe/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores seteo mi puerto, el Handler y pongo a aescuchar al server
func Manejadores() {
	router := mux.NewRouter() //captura el https

	//Enpoints.
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		//si no hay nada le seteamos.
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	//pone al servidor a escuchar el puerto
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
