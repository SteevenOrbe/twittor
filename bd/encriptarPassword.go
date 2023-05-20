package bd

import (
	"golang.org/x/crypto/bcrypt"
)

// EncriptarPassword la passsword 8 pasadas.
func EncriptarPassword(pass string) (string, error) {
	costo := 8 //mayor el costo mayor segurida en la password 8 pasadas.
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
