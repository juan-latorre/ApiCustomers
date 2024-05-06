package server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/Usuarios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
		    ObtenerColeccionUsuariosMongo(w, r)
		case http.MethodPost:
			LlenarColeccionUsuariosMongo(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Metodo no permitido")
			return
		}
	})
}
