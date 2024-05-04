package database

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	crypto "ApiCustomers/crypto"
	m "ApiCustomers/models"

	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "UserAdminDBMongo"
	pwd      = "P4ssW0rd2024_Mer"
	host     = "cluster0.cocts5i.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	database = "Customers"
)

func GetCollection(collection string) *mongo.Collection {

	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", usr, pwd, host)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatalf("Error al conectar con la base de datos: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error al realizar el ping a la base de datos: %v", err)
	}

	db := client.Database(database)
	col := db.Collection(collection)

	return col
}

func SaveCollection() {
	// Conectarse a MongoDB
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", usr, pwd, host)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Colección en la que se insertarán los registros
	collection := client.Database("Customers").Collection("Usuarios")

	// Obtener los datos del endpoint
	resp, err := http.Get("https://62433a7fd126926d0c5d296b.mockapi.io/api/v1/usuarios")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var usuarioList m.ResponseList

	if err := json.NewDecoder(resp.Body).Decode(&usuarioList); err != nil {
		log.Fatal(err)
	}

	var documentos []interface{}
	for _, usuario := range usuarioList {
		//Encriptación de datos antes del guardado en mongoDB
		usuario.CreditCardNum, _ = crypto.EncryptData(usuario.CreditCardNum)
		usuario.CreditCardCcv, _ = crypto.EncryptData(usuario.CreditCardCcv)
		usuario.CuentaNumero, _ = crypto.EncryptData(usuario.CuentaNumero)
		//...
		//Fin
		documentos = append(documentos, usuario)
	}

	// Insertar cada usuario en la colección de MongoDB
	if _, err := collection.InsertMany(context.Background(), documentos); err != nil {
		log.Fatal(err)
	}

}
