package customer_repo

import (
	"ApiCustomers/database"
	m "ApiCustomers/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("Usuarios")
var ctx = context.Background()

func ObtenerUsuariosLista() (m.ResponseList, error) {

	var customersList m.ResponseList

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)
    
	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var customer m.Response
		err = cur.Decode(&customer)

		if err != nil {
			return nil, err
		}

		customersList = append(customersList, &customer)
	}

	return customersList, nil
}

func GuardarUsuariosLista() {
	database.SaveCollection()
}
