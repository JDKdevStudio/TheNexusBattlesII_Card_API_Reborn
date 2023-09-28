package models

import (
	"TheNexusBattlesII_Card_API_Reborn/utils"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConsumibleModel struct {
	Id          primitive.ObjectID     `bson:"_id,omitempty"`
	Imagen      *string                `bson:"imagen,omitempty"                          validate:"required"`
	Icono       *string                `bson:"icono,omitempty"        form:"icono"       validate:"required,url"`
	Nombre      *string                `bson:"nombre,omitempty"       form:"nombre"      validate:"required,regexp=^[A-Za-zÀ-ÿ][A-Za-zÀ-ÿ ]*$"`
	Clase       *string                `bson:"clase,omitempty"        form:"clase"       validate:"required,alpha,oneof=Guerrero Mago Pícaro"`
	Tipo        *string                `bson:"tipo,omitempty"         form:"tipo"        validate:"required,alpha,oneof=Tanque Armas Fuego Hielo Veneno Machete"`
	Coleccion   *string                `bson:"coleccion,omitempty"                       validate:"required,alpha,oneof=Heroes Armas Armaduras Items Epicas"`
	Efecto      map[string]interface{} `bson:"efecto,omitempty"       form:"efecto"      validate:"required"`
	EfectoHeroe map[string]interface{} `bson:"efectoHeroe,omitempty"  form:"efectoHeroe" validate:"required"`
	Descripcion *string                `bson:"descripcion,omitempty"  form:"descripcion" validate:"required,regexp=^[A-Za-zÀ-ÿ][A-Za-zÀ-ÿ ]*$"`
	Precio      *int                   `bson:"precio,omitempty"       form:"precio"      validate:"required,gte=0"`
	Descuento   *int                   `bson:"descuento,omitempty"    form:"descuento"   validate:"required,gte=0"`
	Stock       *int                   `bson:"stock,omitempty"        form:"stock"       validate:"required,gte=0"`
	Estado      *bool                  `bson:"estado,omitempty"       form:"estado"      validate:"required"`
}

func (h *ConsumibleModel) Validate(validateNotNulls bool) error {
	validate := validator.New()
	validate.RegisterValidation("regexp", utils.RegexValidator())
	if validateNotNulls {
		return utils.NullCheckValidator(h, validate)
	}
	err := validate.Struct(h)
	if err != nil {
		return err
	}
	return nil
}
