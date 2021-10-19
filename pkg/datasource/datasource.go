package datasource

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Datasource interface {
	ExecuteQuery(query bson.M) ([]bson.M, error)
	ExecuteAggregate(stages mongo.Pipeline) ([]bson.M, error)
	InsertDocument(document interface{}) ( error)
	DeleteDocument(query bson.M) ( error)
	UpdateDocument(query bson.M, document interface{}) (error)
}

type datasource struct {
	mongoDts MongoDatasource
}

type MongoFields struct {
	Key string `json:"key,omitempty"`
	// ObjectId() or objectid. ObjectID is deprecated--use primitive instead
	ID primitive.ObjectID `bson:"_id, omitempty"`
	// Use these field tags so Golang knows how to map MongoDB fields
	// `bson:"string field" json:"string field"`
	StringField string `bson:"string field" json:"string field"`
	IntField    int    `bson:"int field" json:"int field"`
	BoolField   bool   `bson:"bool field" json:"bool field"`
}

func NewDatasource(mongoDts MongoDatasource) (Datasource, error) {

	return &datasource{mongoDts: mongoDts}, nil
}

func (d *datasource)	InsertDocument(document interface{}) (error){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := d.mongoDts.collection.InsertOne(ctx, document)

	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("could not insert document")
	}
	fmt.Printf("%s",res.InsertedID)
	return nil
}
func (d *datasource)	DeleteDocument(query bson.M) (error){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := d.mongoDts.collection.DeleteOne(ctx, query)

	if err != nil {
		return err
	}
	if res == nil {
		return errors.New("could not delete document")
	}
	fmt.Printf("%d",res.DeletedCount)
	return nil
}
func (d *datasource)	UpdateDocument(query bson.M, document interface{}) (error){
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	res, err := d.mongoDts.collection.UpdateOne(ctx, query, document)

	if err != nil {
		return err
	}
	if res.UpsertedID == nil || res.ModifiedCount == 0 {
		return errors.New("could not update document")
	}
	fmt.Printf("%s",res.UpsertedID)
	return nil
}

func (d *datasource) ConvertQueryToBson(query string) (bson.M, error) {
	// Declare an empty BSON Map object
	var bsonMap bson.M

	// Use the JSON package's Unmarshal() method
	err := json.Unmarshal([]byte(query), &bsonMap)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
		return nil, err
	} else {
		fmt.Println("bsonMap:", bsonMap)
		fmt.Println("bsonMap TYPE:", reflect.TypeOf(bsonMap))
		fmt.Println("BSON:", reflect.TypeOf(bson.M{"int field": bson.M{"$gt": 42}}))
	}
	return bsonMap, nil
}
func (d *datasource) ExecuteQuery(query bson.M) ([]bson.M, error) {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	infoCursor, err := d.mongoDts.collection.Find(ctx, query)

	if err != nil {
		log.Printf("Error executing mongo collection aggregate %s", err.Error())
		return nil, err
	}

	var data []bson.M
	if err = infoCursor.All(ctx, &data); err != nil {
		log.Printf("Error err mongo cursor %s", err.Error())
		return nil, err
	}

	return data, nil
}

func (d *datasource) ExecuteAggregate(stages mongo.Pipeline) ([]bson.M, error) {

	log.Printf("\nStages for Aggregation Framework \n%v\n\n", stages)

	ctx, _ := context.WithTimeout(context.Background(), 1*time.Minute)
	infoCursor, err := d.mongoDts.collection.Aggregate(ctx, stages)

	if err != nil {
		log.Printf("Error executing mongo collection aggregate %v", err)
		return nil, err
	}

	var data []bson.M
	if err = infoCursor.All(ctx, &data); err != nil {
		log.Printf("Error err mongo cursor %v", err)
		return nil, err
	}

	return data, nil
}

// ConvertQueryToPipeline gets aggregation pipeline from a string
func (d *datasource) ConvertQueryToPipeline(str string) mongo.Pipeline {
	log.Printf("Query to convert to pipeline %s", str)
	var pipeline = []bson.D{}
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var doc bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &doc)
		pipeline = append(pipeline, doc)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}
