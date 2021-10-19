package datasource

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Database
var MongoClient *mongo.Client
var ctx = context.TODO()

type MongoConnect struct {
	URI      string
	Database string
}

func NewDBClientMong(m MongoConnect) *mongo.Client {
	clientOptions := options.Client().ApplyURI(m.URI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Connection with Mongo db...OK")
	return client
}

func (c *MongoConnect) GetMongoDB(m *mongo.Client) *mongo.Database {
	MongoDB = m.Database(c.Database)

	return MongoDB
}

func (c *MongoConnect) DisconnectDBMongo(ctx context.Context, m *mongo.Client) {

	err := m.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
