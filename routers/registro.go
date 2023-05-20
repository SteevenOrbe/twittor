package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SteevenOrbe/twittor/bd"
	"github.com/SteevenOrbe/twittor/models"
)

// Registro es la func para crear en la BD el registro de usuario => esto es un METODO ya que no retornara nada.
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos: "+err.Error(), 400)
		return
	}

	//Validaciones
	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido:", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password es menor a 6 caracteres:", 400)
		return
	}

	//Validacion contra los datos.
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya Existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un Error al intentar realizar el registro de usuario", 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logro insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
