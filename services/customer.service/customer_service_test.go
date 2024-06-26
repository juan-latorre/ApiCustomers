package customer_service_test

import (
	customer_service "ApiCustomers/services/customer.service"
	"testing"
)

func ObtenerTodosUsuarios(t *testing.T) {

	customers, err := customer_service.ObtenerColleccionUsuarios() //from customer_service.go

	if err != nil {
		t.Error("Se ha presentado un error en la consulta de clientes")
		t.Fail()
	}

	if len(customers) == 0 {
		t.Error("La consulta no retorno datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizao con exito!")
	}

}
