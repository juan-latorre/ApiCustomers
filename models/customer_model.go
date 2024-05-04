package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Datos del Cliente
type Customer struct {
	ID                        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FecAlta                   string             `bson:"fec_alta" json:"fec_alta"`
	UserName                  string             `bson:"user_name" json:"user_name"`
	CodigoZip                 string             `bson:"codigo_zip" json:"codigo_zip"`
	CreditCardNum             string             `bson:"credit_card_num" json:"credit_card_num"`
	CreditCardCcv             string             `bson:"credit_card_ccv" json:"credit_card_ccv"`
	CuentaNumero              string             `bson:"cuenta_numero" json:"cuenta_numero"`
	Direccion                 string             `bson:"direccion" json:"direccion"`
	GeoLatitud                string             `bson:"geo_latitud" json:"geo_latitud"`
	GeoLongitud               string             `bson:"geo_longitud" json:"geo_longitud"`
	ColorFavorito             string             `bson:"color_favorito" json:"color_favorito"`
	FotoDni                   string             `bson:"foto_dni" json:"foto_dni"`
	IP                        string             `bson:"ip" json:"ip"`
	Auto                      string             `bson:"auto" json:"auto"`
	AutoModelo                string             `bson:"auto_modelo" json:"auto_modelo"`
	AutoTipo                  string             `bson:"auto_tipo" json:"auto_tipo"`
	AutoColor                 string             `bson:"auto_color" json:"auto_color"`
	CantidadComprasRealizadas int                `bson:"cantidad_compras_realizadas" json:"cantidad_compras_realizadas"`
	Avatar                    string             `bson:"avatar" json:"avatar"`
	FecBirthday               time.Time          `bson:"fec_birthday" json:"fec_birthday"`
	Reg                       string             `bson:"registro" json:"registro"`
}

// List of Customers
type CustomersList []*Customer
