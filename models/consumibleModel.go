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
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConsumibleModel struct {
	Id          primitive.ObjectID     `bson:"_id,omitempty"`
	Imagen      *string                `bson:"imagen,omitempty"                          validate:"required"`
	Icono       *string                `bson:"icono,omitempty"        form:"icono"       validate:"required,url"`
	Nombre      *string                `bson:"nombre,omitempty"       form:"nombre"      validate:"required,regexp=^[A-Za-zÀ-ÿ][A-Za-zÀ-ÿ ]*$"`
	Clase       *string                `bson:"clase,omitempty"        form:"clase"       validate:"required,alpha,oneof=Guerrero Mago Pícaro"`
	Tipo        *string                `bson:"tipo,omitempty"         form:"tipo"        validate:"required,alpha,oneof=Tanque Armas Fuego Hielo Veneno Machete"`
	Coleccion   *string                `bson:"coleccion,omitempty"                       validate:"required,alpha,oneof=Heroes Armas Armaduras Items Épicas"`
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

func (ConsumibleModel) GetObjectList(query PaginationModel, coleccion string, consumibleList *[]ConsumibleModel) error {
	filter := bson.M{"coleccion": coleccion}
	if *query.StatusFilter {
		filter["estado"] = query.StatusFilter
	}
	return getConsumibleListRaw(filter, query, consumibleList)
}

func (ConsumibleModel) GetObjectListByName(query PaginationModel, keyword string, coleccion string, consumibleList *[]ConsumibleModel) error {
	filter := bson.M{"nombre": bson.M{"$regex": keyword, "$options": "i"}, "estado": true, "coleccion": coleccion}
	return getConsumibleListRaw(filter, query, consumibleList)
}

// [Funciones Raw]
func getConsumibleListRaw(filter interface{}, query PaginationModel, consumibleList *[]ConsumibleModel) error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	response, err := col.Find(ctx, filter, options.Find().SetSkip((int64(query.Page-1) * int64(query.Size))).SetLimit(int64(query.Size)))
	if err != nil {
		return err
	}
	defer response.Close(ctx)
	*consumibleList = []ConsumibleModel{}
	for response.Next(ctx) {
		var element ConsumibleModel
		if err := response.Decode(&element); err != nil {
			return err
		}
		*consumibleList = append(*consumibleList, element)
	}
	return nil
}
