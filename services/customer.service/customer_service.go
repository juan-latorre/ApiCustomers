package customer_service

import (
	custormer_repositorio "ApiCustomers/Repositories/customer.repo"
	m "ApiCustomers/models"
)

func GetInfoCustormer() (m.ResponseList, error) {

	customers, err := custormer_repositorio.GetBasicInfo()

	if err != nil {
		return nil, err
	}

	return customers, nil
}

func FullFillCollectionFromEndPoint() {

	custormer_repositorio.SaveInfoIntoCollectionDB()

}
