package server

import (
	"net/http"
	m "ApiCustomers/models"

)

var listaUsuarios m.ResponseList

func New(addr string) *http.Server {
	initRoutes()
	return &http.Server{
		Addr: addr,
	}
}