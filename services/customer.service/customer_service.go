package customer_service

import (
	custormer_repositorio "ApiCustomers/Repositories/customer.repo"
	m "ApiCustomers/models"
)

func GetInfoCustormer() (m.CustomersList, error) {

	customers, err := custormer_repositorio.GetBasicInfo()

	if err != nil {
		return nil, err
	}

	return customers, nil
}


