package datasource

import (
	"log"

	"github.com/hellofresh/janus/pkg/plugin/basic/encrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDatasource represents a mongodb datasource
type MongoDatasource struct {
	collection *mongo.Collection
	hash       encrypt.Hash
}

// NewMongoDatasource creates a mongo API definition dts
func NewMongoDatasource(db *mongo.Database, collectionName string) (*MongoDatasource, error) {
	log.Print("Init mongo collection ", collectionName)
	return &MongoDatasource{collection: db.Collection(collectionName), hash: encrypt.Hash{}}, nil
}
