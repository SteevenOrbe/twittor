package bd

import (
	"github.com/SteevenOrbe/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

// IntengoLogin realiza el chequeo de login a la BD.
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	user, encontrado, _ := ChequeoYaExisteUsuario(email)
	if encontrado == false {
		return user, false
	}
	passwordBytes := []byte(password)   //la que me vino en parametros.
	passwordBD := []byte(user.Password) //passwor que eta grabada en la BD, esta encritada.

	//Desencritar la pass-> primero va la encritada y luego l sin encritar.
	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
