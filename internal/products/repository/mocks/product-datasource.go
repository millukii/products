package mocks

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

type FakeDatasource struct {
	MockExecuteQuery func(query bson.M) ([]bson.M, error)
	MockExecuteAggregate func(stages mongo.Pipeline) ([]bson.M, error)
	MockInsertDocument func(document interface{}) error
	MockDeleteDocument func(query bson.M) error
	MockUpdateDocument func(query bson.M, document interface{}) error
}

func (f FakeDatasource) ExecuteQuery(query bson.M) ([]bson.M, error) {
	return f.MockExecuteQuery(query)
}

func (f FakeDatasource) ExecuteAggregate(stages mongo.Pipeline) ([]bson.M, error) {
	return f.MockExecuteAggregate(stages)
}
func (f FakeDatasource) InsertDocument(document interface{}) error {
	return f.MockInsertDocument(document)
}
func (f FakeDatasource) DeleteDocument(query bson.M) error {
	return f.MockDeleteDocument(query)
}
func (f FakeDatasource) UpdateDocument(query bson.M, document interface{}) error {
	return f.MockUpdateDocument(query,document)
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
