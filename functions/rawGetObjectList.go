package functions

import (
	"TheNexusBattlesII_Card_API_Reborn/database"
	"TheNexusBattlesII_Card_API_Reborn/models"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
)

func rawGetObjectList(filter interface{}, query models.PaginationModel, list *[]map[string]interface{}) error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	response, err := col.Find(ctx, filter, options.Find().SetSkip((int64(query.Page-1) * int64(query.Size))).SetLimit(int64(query.Size)))
	if err != nil {
		return err
	}
	defer response.Close(ctx)
	for response.Next(ctx) {
		var document = make(map[string]interface{})
		if err := response.Decode(document); err != nil {
			return err
		}
		*list = append(*list, document)
	}
	return nil
}
