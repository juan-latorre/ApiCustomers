package main

import (
	crypto "ApiCustomers/crypto"
	customer_service "ApiCustomers/services/customer.service"
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Listado Clientes")

	//customer_service.FullFillCollectionFromEndPoint()

	//os.Exit(1)

	formatDate := "02/01/2006"

	customers, err := customer_service.GetInfoCustormer()

	for _, customer := range customers {
		//fmt.Println(customer)
		//fmt.Printf("Numero: %s \n",customer.Reg)

		// encryptedData, err := crypto.EncryptData(customer.CreditCardNum)
		// if err != nil {
		// 	fmt.Println("Error al cifrar:", err)
		// 	return
		// }

		// Descifrar el atributo
		decryptedData, err := crypto.DecryptData(customer.CreditCardNum)
		if err != nil {
			fmt.Println("Error al descifrar:", err)
			return
		}

		fmt.Printf("Id: %s Nombre: %s CuentaNo. %s CreditCard Decrypted: %s ccv: %s Fecha: %s \n",
			customer.ID,
			customer.UserName,
			customer.CuentaNumero,
			decryptedData,
			customer.CreditCardCcv,
			customer.FecBirthday.Format(formatDate))
	}

	if err != nil {
		fmt.Println(err.Error())
	}

}
