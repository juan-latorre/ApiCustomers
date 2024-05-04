package customer_repo

import (
	"ApiCustomers/database"
	m "ApiCustomers/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("People")
var ctx = context.Background()

func GetBasicInfo() (m.CustomersList, error) {

	var customersList m.CustomersList

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var customer m.Customer
		err = cur.Decode(&customer)

		if err != nil {
			return nil, err
		}

		customersList = append(customersList, &customer)
	}

	return customersList, nil
}
