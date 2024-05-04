package main

import (
	crypto "ApiCustomers/crypto"
	customer_service "ApiCustomers/services/customer.service"
	"fmt"
	//"time"
)

func main() {
	fmt.Println("Listado Clientes")

	formatDate := "02/01/2006"

	customers, err := customer_service.GetInfoCustormer()

	for _, customer := range customers {
		//fmt.Println(customer)
		//fmt.Printf("Numero: %s \n",customer.Reg)

		encryptedData, err := crypto.EncryptData(customer.CreditCardNum)
		if err != nil {
			fmt.Println("Error al cifrar:", err)
			return
		}

		// Descifrar el atributo
		decryptedData, err := crypto.DecryptData(encryptedData)
		if err != nil {
			fmt.Println("Error al descifrar:", err)
			return
		}

		//customer.CreditCardNum = encryptedData

		fmt.Printf("Registro: %s Nombre: %s CuentaNo. %s CreditCard Decrypted: %s CreditCard Encrypted: %s ccv: %s Fecha: %s \n",
			customer.Reg,
			customer.UserName,
			customer.CuentaNumero,
			decryptedData,
			encryptedData,
			customer.CreditCardCcv,
			customer.FecBirthday.Format(formatDate))
	}

	if err != nil {
		fmt.Println(err.Error())
	}

}
