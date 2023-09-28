package functions

import (
	"TheNexusBattlesII_Card_API_Reborn/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetObjectList(query models.PaginationModel, list *[]map[string]interface{}) error {
	filter := bson.M{"nombre": bson.M{"$regex": query.Keyword, "$options": "i"}}
	if query.Coleccion != "All" {
		filter["coleccion"] = query.Coleccion
	}
	if *query.OnlyActives {
		filter["estado"] = query.OnlyActives
	}
	return rawGetObjectList(filter, query, list)
}
