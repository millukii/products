package mocks

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

type FakeDatasource struct {
}

func (f FakeDatasource) ExecuteQuery(query bson.M) ([]bson.M, error) {
	return []bson.M{
		bson.M{"sku": "1", "name": "rasho laser", "brand": "monster", "size": "M", "price": 200, "principalImage": "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ",
			"otherImages": bson.A{"https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ", "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ"},
		},
	}, nil
}

func (f FakeDatasource) ExecuteAggregate(stages mongo.Pipeline) ([]bson.M, error) {
	return []bson.M{
		bson.M{"sku": "1", "name": "rasho laser", "brand": "monster", "size": "M", "price": 200, "principalImage": "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ",
			"otherImages": bson.A{"https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ", "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ"},
		},
		bson.M{"sku": "2", "name": "rasho laser2", "brand": "acme", "size": "XXL", "price": 900, "principalImage": "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ",
			"otherImages": bson.A{"https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ", "https://i.picsum.photos/id/1/5616/3744.jpg?hmac=kKHwwU8s46oNettHKwJ24qOlIAsWN9d2TtsXDoCWWsQ"},
		},
	}, nil
}
func (f FakeDatasource) InsertDocument(document interface{}) error {
	return nil
}
func (f FakeDatasource) DeleteDocument(query bson.M) error {
	return nil
}
func (f FakeDatasource) UpdateDocument(query bson.M, document interface{}) error {
	return nil
}


type FakeDatasourceWithProblems struct {
}

func (f FakeDatasourceWithProblems) ExecuteQuery(query bson.M) ([]bson.M, error) {
	return nil, errors.New("Mock error")
}

func (f FakeDatasourceWithProblems) ExecuteAggregate(stages mongo.Pipeline) ([]bson.M, error) {
	return nil, errors.New("Mock error")
}
func (f FakeDatasourceWithProblems) InsertDocument(document interface{}) error {
	return errors.New("Mock error")
}
func (f FakeDatasourceWithProblems) DeleteDocument(query bson.M) error {
	return errors.New("Mock error")
}
func (f FakeDatasourceWithProblems) UpdateDocument(query bson.M, document interface{}) error {
	return errors.New("Mock error")
}