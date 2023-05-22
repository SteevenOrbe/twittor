package routers

import (
	"errors"
	"strings"

	"github.com/SteevenOrbe/twittor/bd"
	"github.com/SteevenOrbe/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email valor email usado en todos los EndPoint.
var Email string

// IDUuario es el ID devuelto del modelo, que se usara en todos los enpoint.
var IDUuario string

// ProcesoToken proceso token para extraer sus valore.
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("developer-kushki")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	//Quita los espacio que pueda tener el token.
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil //miClave lo conbierte en una interface
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Yoken Invalido")
	}
	return claims, false, string(""), err
}
