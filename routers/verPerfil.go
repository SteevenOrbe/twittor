package routers

import (
	"encoding/json"
	"net/http"

	"github.com/SteevenOrbe/twittor/bd"
)

// VerPerfil permite extraer los valores del perfil.
func Verperfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	// y si encuentra el perfil.
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(perfil)

}
