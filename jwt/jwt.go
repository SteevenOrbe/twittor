package jwt

import (
	"time"

	"github.com/SteevenOrbe/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GeneroJWT genera el encriptado con jwt.
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("developer-kushki")

	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografoa":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioweb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
