package functions

import (
	"TheNexusBattlesII_Card_API_Reborn/database"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObject(item_id primitive.ObjectID, item map[string]interface{}) error {
	db := database.GetMongoClient()
	col := db.Database(os.Getenv("MONGO_DB")).Collection("cartas")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return col.FindOne(ctx, bson.M{"_id": item_id}).Decode(item)
}
