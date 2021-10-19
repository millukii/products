package repository_test

import (
	"api-products/internal/products/models"
	"api-products/internal/products/repository"
	"api-products/internal/products/repository/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestProductRepository(t *testing.T) {

	//dependencies

 	dts := &mocks.FakeDatasource{}
	repo,_ := repository.NewProductsRepository(dts)

	dtsWithProblems := &mocks.FakeDatasourceWithProblems{}

	repoErr,_ := repository.NewProductsRepository(dtsWithProblems)

	ctx := context.Background()

	t.Run("GetAll", func(t *testing.T) {


		t.Run("Datasource respond data without errors", func(t *testing.T){
			products, err := repo.GetAll(ctx)
			
			assert.NoError(t, err)
			assert.NotNil(t, products)
		})
		t.Run("Datasource respond errors instead of the product list", func(t *testing.T){
			products, err := repoErr.GetAll(ctx)
			
			assert.Error(t, err)
			assert.Nil(t, products)
		})
	})
	t.Run("GetBy", func(t *testing.T) {

		t.Run("Datasource respond data of one product without errors", func(t *testing.T){
			product, err := repo.GetBy(ctx, "sku", "1")
			
			assert.NoError(t, err)
			assert.NotNil(t, product)
		})
		t.Run("Datasource respond errors trying to get one product", func(t *testing.T){
			product, err := repoErr.GetBy(ctx, "sku", "1")
			
			assert.Error(t, err)
			assert.Nil(t, product)
		})
	})
	t.Run("Add", func(t *testing.T) {
		t.Run("Datasource insert one product without errors", func(t *testing.T){
			product, err := repo.Add(ctx, &models.Product{
				SKU: "3",
				Name: "jojo",
				Brand: "jojo2",
				Size: "12",
				Price: 10.0,
				PrincipalImage: "",
				OtherImages: []string{"", ""},
			})
			
			assert.NoError(t, err)
			assert.NotNil(t, product)
		})
		t.Run("Datasource cannot insert the product", func(t *testing.T){
			product, err := repoErr.Add(ctx, &models.Product{
				SKU: "3",
				Name: "jojo",
				Brand: "jojo2",
				Size: "12",
				Price: 10.0,
				PrincipalImage: "",
				OtherImages: []string{"", ""},
			})
			
			assert.Error(t, err)
			assert.Nil(t, product)
		})
	})
	t.Run("Delete", func(t *testing.T) {
		t.Run("Datasource delete one product without errors", func(t *testing.T){
			err := repo.Delete(ctx,"1")
			
			assert.NoError(t, err)
		})
		t.Run("Datasource cannot delete the product", func(t *testing.T){
			err := repoErr.Delete(ctx,"1")
			
			assert.Error(t, err)
		})
	})
	t.Run("Update", func(t *testing.T) {
		t.Run("Datasource update one product without errors", func(t *testing.T){
			product, err := repo.Update(ctx, "1", &models.Product{})
			
			assert.NoError(t, err)
			assert.NotNil(t, product)
		})
		t.Run("Datasource cannot update the product", func(t *testing.T){
			product, err := repoErr.Update(ctx, "1", &models.Product{})
			
			assert.Error(t, err)
			assert.Nil(t, product)
		})
	})
}