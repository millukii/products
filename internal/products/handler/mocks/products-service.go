package mocks

import (
	"api-products/internal/products/models"
	"context"
)

type ServiceMock struct {
	MockGetBySKU func(ctx context.Context, sku string) (product *models.Product, err error)
	MockGetProducts func(ctx context.Context) ([]*models.Product, error)
	MockDeleteProductBySKU func(ctx context.Context, sku string) error
	MockNewProduct func(ctx context.Context, product *models.Product) (*models.Product, error)
}

func (s ServiceMock) GetBySKU(ctx context.Context, sku string) (*models.Product, error) {
	return s.MockGetBySKU(ctx, sku)
}
func (s ServiceMock) GetProducts(ctx context.Context) ([]*models.Product, error) {
	return s.MockGetProducts(ctx)
}
func (s ServiceMock) DeleteProductBySKU(ctx context.Context, sku string) error {
	return s.MockDeleteProductBySKU(ctx, sku)
}
func (s ServiceMock) UpdateProductBySKU(ctx context.Context, sku string, product *models.Product) (*models.Product, error) {
	return nil, nil
}
func (s ServiceMock) NewProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	return s.MockNewProduct(ctx,product)
}