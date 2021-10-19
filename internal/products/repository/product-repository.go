package repository

import (
	"api-products/internal/products/models"
	"api-products/pkg/datasource"
	"context"
	"encoding/json"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductsRepository interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetBy(ctx context.Context, column, value string) (*models.Product, error)
	Add(ctx context.Context, product *models.Product) (*models.Product, error)
	Delete(ctx context.Context, sku string) (error)
	Update(ctx context.Context, sku string, product *models.Product) (*models.Product, error)
}

type repository struct {
	datasource datasource.Datasource
}

func NewProductsRepository(dts datasource.Datasource) (ProductsRepository, error) {
	return &repository{datasource: dts}, nil
}

func (r repository) 	GetBy(ctx context.Context, column, value string) (*models.Product, error) {

	data := make([]*models.Product, 0)

	matchStage := bson.D{{"$match", bson.D{{column, value}}}}

	docs, err := r.datasource.ExecuteAggregate(mongo.Pipeline{matchStage})
	if err != nil {
		return nil, err
	}

	raw, _ := json.Marshal(docs)
	json.Unmarshal(raw, &data)

	log.Println("Consulta desde mongo recuperada correctamente")
	return data[0], nil
}
func (r repository) 		GetAll(ctx context.Context) ([]*models.Product, error){
	data := make([]*models.Product, 0)
	docs, err := r.datasource.ExecuteQuery(bson.M{})
	if err != nil {
		return nil, err
	}

	raw, _ := json.Marshal(docs)
	//log.Printf("Raw from mongo %s", raw)
	json.Unmarshal(raw, &data)

	log.Println("Consulta desde mongo recuperada correctamente")
	return data, nil
}
func (r repository) 		Add(ctx context.Context, product *models.Product) (*models.Product, error){

	err :=	r.datasource.InsertDocument(product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (r repository) 		Delete(ctx context.Context, sku string) (error){

	err :=	r.datasource.DeleteDocument(bson.M{"sku":sku})
	if err != nil {
		return err
	}
	return nil
}
func (r repository) 		Update(ctx context.Context, sku string, product *models.Product) (*models.Product, error){

	err :=	r.datasource.UpdateDocument(bson.M{"sku":sku},product)
	if err != nil {
		return nil, err
	}
	return product, nil
}