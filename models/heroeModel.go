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

type HeroeModel struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Imagen      *string            `bson:"imagen,omitempty"                          validate:"required"`
	Icono       *string            `bson:"icono,omitempty"        form:"icono"       validate:"required,url"`
	Nombre      *string            `bson:"nombre,omitempty"       form:"nombre"      validate:"required,regexp=^[A-Za-zÀ-ÿ][A-Za-zÀ-ÿ ]*$"`
	Clase       *string            `bson:"clase,omitempty"        form:"clase"       validate:"required,alpha,oneof=Guerrero Mago Pícaro"`
	Tipo        *string            `bson:"tipo,omitempty"         form:"tipo"        validate:"required,alpha,oneof=Tanque Armas Fuego Hielo Veneno Machete"`
	Coleccion   *string            `bson:"coleccion,omitempty"                       validate:"required,alpha,oneof=Heroes Armas Armaduras Items Épicas"`
	Poder       *int               `bson:"poder,omitempty"        form:"poder"       validate:"required,gt=0"`
	Vida        *int               `bson:"vida,omitempty"         form:"vida"        validate:"required,gt=0"`
	Defensa     *int               `bson:"defensa,omitempty"      form:"defensa"     validate:"required,gt=0"`
	AtaqueBase  *int               `bson:"ataqueBase,omitempty"   form:"ataqueBase"  validate:"required,gt=0"`
	AtaqueRnd   *int               `bson:"ataqueRnd,omitempty"    form:"ataqueRnd"   validate:"required,gt=0"`
	Dano        *int               `bson:"daño,omitempty"         form:"daño"        validate:"required,gt=0"`
	Descripcion *string            `bson:"descripcion,omitempty"  form:"descripcion" validate:"required,regexp=^[A-Za-zÀ-ÿ][A-Za-zÀ-ÿ ]*$"`
	Precio      *int               `bson:"precio,omitempty"       form:"precio"      validate:"required,gte=0"`
	Descuento   *int               `bson:"descuento,omitempty"    form:"descuento"   validate:"required,gte=0"`
	Stock       *int               `bson:"stock,omitempty"        form:"stock"       validate:"required,gte=0"`
	Estado      *bool              `bson:"estado,omitempty"       form:"estado"      validate:"required"`
}

func (h *HeroeModel) Validate(validateNotNulls bool) error {
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

func (h *HeroeModel) GetObject(heroe_id primitive.ObjectID) error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return col.FindOne(ctx, bson.M{"_id": heroe_id, "estado": true}).Decode(&h)
}

func (HeroeModel) GetObjectList(query PaginationModel, heroeList *[]HeroeModel) error {
	filter := bson.M{"coleccion": "Heroes"}
	if *query.StatusFilter {
		filter["estado"] = query.StatusFilter
	}
	return getObjectListRaw(filter, query, heroeList)
}

func (HeroeModel) GetObjectListByName(query PaginationModel, keyword string, heroeList *[]HeroeModel) error {
	filter := bson.M{"nombre": bson.M{"$regex": keyword, "$options": "i"}, "estado": true, "coleccion": "Heroes"}
	return getObjectListRaw(filter, query, heroeList)
}

func (h HeroeModel) PostObject() error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if _, err := col.InsertOne(ctx, h); err != nil {
		return err
	}
	return nil
}

func (h *HeroeModel) UpdateObject() error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	var updated_heroe HeroeModel
	if err := col.FindOneAndUpdate(ctx, bson.M{"_id": h.Id}, bson.M{"$set": h}).Decode(&updated_heroe); err != nil {
		return err
	}
	if h.Imagen != nil {
		go utils.DeleteImageHandler(updated_heroe.Imagen)
	}
	return nil
}

// [Funciones Raw]
func getObjectListRaw(filter interface{}, query PaginationModel, heroeList *[]HeroeModel) error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	response, err := col.Find(ctx, filter, options.Find().SetSkip((int64(query.Page-1) * int64(query.Size))).SetLimit(int64(query.Size)))
	if err != nil {
		return err
	}
	defer response.Close(ctx)
	*heroeList = []HeroeModel{}
	for response.Next(ctx) {
		var element HeroeModel
		if err := response.Decode(&element); err != nil {
			return err
		}
		*heroeList = append(*heroeList, element)
	}
	return nil
}
