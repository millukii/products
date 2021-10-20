package service_test

import (
	"api-products/internal/products/models"
	"api-products/internal/products/service"
	"api-products/internal/products/service/mocks"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

)

func TestService(t *testing.T) {

	repo := &mocks.RepositoryMock{}
	ctx := context.Background()
	t.Run("GetProducts", func(t *testing.T) {

		t.Run("Success", func(t *testing.T) {
			repo.MockGetAll=respMockGetAll
			svc := service.NewProductService(repo)
			products, err := svc.GetProducts(ctx)

			assert.Nil(t, err)
			assert.NotNil(t, products)
		})
	})
		t.Run("GetBySKU", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				repo.MockGetBy=respMockGetBy
				svc := service.NewProductService(repo)
				product, err :=svc.GetBySKU(ctx, "1")
				assert.Nil(t, err)
				assert.NotNil(t, product)
			})
	})
	t.Run("DeleteProductBySKU", func(t *testing.T) {
		t.Run("Succes", func(t *testing.T){
			repo.MockDelete=respMockDelete
			svc := service.NewProductService(repo)
			err :=	svc.DeleteProductBySKU(ctx, "1")
			assert.Nil(t, err)
		})
	})
	t.Run("UpdateProductBySKU", func(t *testing.T) {
		t.Run("Success", func(t *testing.T){
			repo.MockUpdate=respMockUpdate
			svc := service.NewProductService(repo)
			product, err :=svc.UpdateProductBySKU(ctx, "1", &models.Product{})
			assert.Nil(t, err)
			assert.NotNil(t, product)
		})
	})
	t.Run("NewProduct", func(t *testing.T) {
		t.Run("Success", func(t *testing.T){
			repo.MockAdd=respMockAdd
			svc := service.NewProductService(repo)
			product,err :=svc.NewProduct(ctx, &models.Product{})
			assert.Nil(t, err)
			assert.NotNil(t, product)
		})
	})
}

func respMockGetAll(ctx context.Context) ([]*models.Product, error){
	var products []*models.Product

	products = append(products, &models.Product{
			SKU: "3",
			Name: "jojo",
			Brand: "jojo2",
			Size: "1",
			Price: 1.0,
			PrincipalImage: "",
			OtherImages: []string{"",""}, 
		})
	return products,nil
}

func	respMockGetBy (ctx context.Context, column, value string) (*models.Product, error){
	return &models.Product{

	}, nil
}
func	respMockAdd (ctx context.Context, product *models.Product) (*models.Product, error){
	return &models.Product{}, nil
}
func	respMockDelete (ctx context.Context, sku string) error{
	return nil
}
func	respMockUpdate (ctx context.Context, sku string, product *models.Product) (*models.Product, error){
	return &models.Product{

	}, nil
}