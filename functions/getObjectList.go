package functions

import (
	"TheNexusBattlesII_Card_API_Reborn/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetObjectList(query models.PaginationModel, list *[]map[string]interface{}) error {
	filter := bson.M{"nombre": bson.M{"$regex": query.Keyword, "$options": "i"}, "estado": true, "coleccion": query.Coleccion}
	return rawGetObjectList(filter, query, list)
}
