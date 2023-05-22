package middlew

import (
	"net/http"

	"github.com/SteevenOrbe/twittor/routers"
)

// ValidoJWT permite validar el JWT que nos viene en la peticion.
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//verifica que sea un token valido (ProcesoToken) me da un string.
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization")) // si el token es valido o no.
		if err != nil {
			http.Error(w, "Error en el Token.! "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
