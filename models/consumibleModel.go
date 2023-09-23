package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConsumibleModle struct {
	Id          primitive.ObjectID
	Imagen      string
	Icono       string
	Nombre      string
	Clase       string
	Tipo        string
	Coleccion   string
	Efecto      map[string]interface{}
	EfectoHeroe map[string]interface{}
	Descripcion string
	Precio      int
	Descuento   int
	Stock       int
	Estado      bool
}
