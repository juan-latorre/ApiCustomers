package customer_service

import (
	custormer_repositorio "ApiCustomers/Repositories/customer.repo"
	m "ApiCustomers/models"
)

func ObtenerColleccionUsuarios() (m.ResponseList, error) {

	customers, err := custormer_repositorio.ObtenerUsuariosLista()

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func LlenarColeccionUsuarios() {

	custormer_repositorio.GuardarUsuariosLista()

}
