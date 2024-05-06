package server

import (
	customer_service "ApiCustomers/services/customer.service"
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Metodo no permitido")
		return
	}

	fmt.Fprintf(w, "OK %s", "visitor")
}

func ObtenerColeccionUsuariosMongo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	listaUsuarios, _ = customer_service.ObtenerColleccionUsuarios()
	json.NewEncoder(w).Encode(listaUsuarios)

}

func LlenarColeccionUsuariosMongo(w http.ResponseWriter, r *http.Request) {

	usr := &listaUsuarios

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	listaUsuarios = append(listaUsuarios, *usr...)
	fmt.Fprintf(w, "La coleccion de usuarios fue llenada")
}
