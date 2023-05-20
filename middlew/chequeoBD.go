package middlew

import (
	"net/http"

	"github.com/SteevenOrbe/twittor/bd"
)

// ChequeoBD es el middler que me perite conocer el estado de la BD.
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		//SI NO FALLA, LE PASO TODOS LOS VALORE RECIVIDOS.
		next.ServeHTTP(w, r)
	}

}
