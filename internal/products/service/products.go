package service

import (
	"api-products/internal/products/models"
	"api-products/internal/products/repository"
	"context"
	"log"
)

type ProductsService interface {
	GetBySKU(ctx context.Context, sku string) (*models.Product, error)
	GetProducts(ctx context.Context) ([]*models.Product, error)
	DeleteProductBySKU(ctx context.Context, sku string) (error)
	UpdateProductBySKU(ctx context.Context, sku string, product *models.Product ) (*models.Product, error)
	NewProduct(ctx context.Context, product *models.Product) (*models.Product, error)
}

type service struct {
	repo repository.ProductsRepository
}

func NewProductService(repo repository.ProductsRepository) ProductsService {

	return &service{repo: repo}
}

func (s service) GetProducts(ctx context.Context) ([]*models.Product, error) {

	log.Println("[service /internal/products/service] GetProducts")

	products, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}
func (s service) 	GetBySKU(ctx context.Context, sku string) (*models.Product, error){

	product, err := s.repo.GetBy(ctx,"sku", sku)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s service) 	DeleteProductBySKU(ctx context.Context, sku string) (error){

 err := s.repo.Delete(ctx, sku)
	if err != nil {
		return  err
	}
	return nil
}
func (s service) 	UpdateProductBySKU(ctx context.Context, sku string, product *models.Product ) (*models.Product, error){
	product, err := s.repo.Update(ctx, sku, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (s service) 	NewProduct(ctx context.Context, product *models.Product) (*models.Product, error){
	product, err := s.repo.Add(ctx,product)
	if err != nil {
		return nil, err
	}
	return product, nil
}