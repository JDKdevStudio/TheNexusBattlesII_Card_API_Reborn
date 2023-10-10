package models

import (
	"TheNexusBattlesII_Card_API_Reborn/database"
	"TheNexusBattlesII_Card_API_Reborn/utils"
	"context"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConsumibleModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Imagen      *string            `bson:"imagen,omitempty"                          validate:"required"`
	Icono       *string            `bson:"icono,omitempty"        form:"icono"       validate:"required,url"`
	Nombre      *string            `bson:"nombre,omitempty"       form:"nombre"      validate:"required"`
	Clase       *string            `bson:"clase,omitempty"        form:"clase"       validate:"required,alpha,oneof=Guerrero Mago Pícaro"`
	Tipo        *string            `bson:"tipo,omitempty"         form:"tipo"        validate:"required,alpha,oneof=Tanque Armas Fuego Hielo Veneno Machete"`
	Coleccion   *string            `bson:"coleccion,omitempty"    form:"coleccion"   validate:"required,alpha,oneof=Armas Armaduras Items Epicas"`
	Efecto      *EfectoModel       `bson:"efecto,omitempty"                          validate:"required"`
	EfectoHeroe *EfectoModel       `bson:"efectoHeroe,omitempty"                     validate:"required"`
	Descripcion *string            `bson:"descripcion,omitempty"  form:"descripcion" validate:"required"`
	Precio      *int               `bson:"precio,omitempty"       form:"precio"      validate:"required,gte=0"`
	Descuento   *int               `bson:"descuento,omitempty"    form:"descuento"   validate:"required,gte=0"`
	Stock       *int               `bson:"stock,omitempty"        form:"stock"       validate:"required,gte=0"`
	Estado      *bool              `bson:"estado,omitempty"       form:"estado"      validate:"boolrequired"`
}

func (h *ConsumibleModel) Validate(validateNotNulls bool) error {
	validate := validator.New()
	validate.RegisterValidation("boolrequired", utils.BoolValidator())
	if validateNotNulls {
		return utils.NullCheckValidator(h, validate)
	}
	err := validate.Struct(h)
	if err != nil {
		return err
	}
	return nil
}

func (h ConsumibleModel) PostObject() error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if _, err := col.InsertOne(ctx, h); err != nil {
		return err
	}
	return nil
}

func (h *ConsumibleModel) UpdateObject() error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var updated_consumible ConsumibleModel
	if err := col.FindOneAndUpdate(ctx, bson.M{"_id": h.Id}, bson.M{"$set": h}).Decode(&updated_consumible); err != nil {
		return err
	}
	if h.Imagen != nil {
		go utils.DeleteImageHandler(updated_consumible.Imagen)
	}
	return nil
}
