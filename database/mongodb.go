package database

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongo_db     *mongo.Client
	syncInstance = &sync.Mutex{}
)

// [database.GetMongoClient()]
// Esta función devuelve un cliente de mongoDB,
// Verifica si está instanciado, y si no crea uno nuevo y bloquea su uso.
func GetMongoClient() *mongo.Client {
	if mongo_db == nil {
		syncInstance.Lock()
		defer syncInstance.Unlock()
		if mongo_db == nil {
			mongo_db = connectMongoClient()
		} else {
			return mongo_db
		}
	} else {
		return mongo_db
	}
	return mongo_db
}

// [database.connectMongoClient]
// Esta función crea la conexión entre mongoDB y la variable del cliente global.
func connectMongoClient() *mongo.Client {
	var mongo_client *mongo.Client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	mongo_client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_CONN")))
	if err != nil {
		log.Fatalf("No se pudo realizar la conexión: %v", err)
	}
	return mongo_client
}
