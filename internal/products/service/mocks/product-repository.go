package mocks

import (
	"api-products/internal/products/models"
	"context"
)

type RepositoryMock struct {
	MockGetAll func(ctx context.Context) ([]*models.Product, error)
	MockGetBy func(ctx context.Context, column, value string) (*models.Product, error)
	MockAdd func(ctx context.Context, product *models.Product) (*models.Product, error)
	MockDelete func(ctx context.Context, sku string) error
	MockUpdate func(ctx context.Context, sku string, product *models.Product) (*models.Product, error)
}

func (m *RepositoryMock) GetAll(ctx context.Context) ([]*models.Product, error){
	return m.MockGetAll(ctx)
}
func (m *RepositoryMock) GetBy(ctx context.Context, column, value string) (*models.Product, error){
	return m.MockGetBy(ctx, column, value)
}
func (m *RepositoryMock) Add(ctx context.Context, product *models.Product) (*models.Product, error){
	return m.MockAdd(ctx,product)
}
func (m *RepositoryMock) Delete(ctx context.Context, sku string) error{
	return m.MockDelete(ctx,sku)
}
func (m *RepositoryMock) Update(ctx context.Context, sku string, product *models.Product) (*models.Product, error){
	return m.MockUpdate(ctx, sku,product)
}