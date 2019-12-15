package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func InitDB() {
	ctx := context.Background()

	client, err := mongo.NewClinet(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func(ctx context.Context, client *mongo.Client) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func checkConnection(ctx context.Context, c *mongo.Client) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return c.Ping(ctx, readpref.Primary())
}
