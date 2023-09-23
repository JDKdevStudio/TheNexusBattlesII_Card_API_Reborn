package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type HeroeModel struct {
	Id          primitive.ObjectID
	Imagen      string
	Icono       string
	Nombre      string
	Clase       string
	Tipo        string
	Coleccion   string
	Poder       int
	Vida        int
	Defensa     int
	AtaqueBase  int
	AtaqueRnd   int
	Daño        int
	Descripcion string
	Precio      int
	Descuento   int
	Stock       int
	Estado      bool
}
