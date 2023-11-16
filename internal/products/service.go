package products

import (
	"context"
	"log"

	"github.com/burgosfacundo/ApiGo.git/internal/domain"
)

// Service represents a contract with all the functions that need to be implemented
type Service interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id string) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id string) (domain.Product, error)
	Delete(ctx context.Context, id string) error
}

// service is a struct that contains the repository of Product objects
type service struct {
	repository Repository
}

// NewServiceProduct is a function that loads the repository into the service
func NewServiceProduct(repository Repository) Service {
	return &service{repository: repository}
}

// Create is a function that calls the repository for create a Product in the db
func (s *service) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	// We call the repository for create a Product
	product, err := s.repository.Create(ctx, product)

	// If we have an error log it and return it
	if err != nil {
		log.Println("[ProductsService][Create] error creating product", err)
		return domain.Product{}, err
	}

	// We return the product
	return product, nil
}

// GetAll is a function that calls the repository for return all the products in the db
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	// We call the repository for get all the products
	listProducts, err := s.repository.GetAll(ctx)

	// If we have an error log it and return it
	if err != nil {
		log.Println("[ProductsService][GetAll] error getting all products", err)
		return []domain.Product{}, err
	}

	// We return the products
	return listProducts, nil
}

// GetById is a function that calls the repository for return a Product by Id
func (s *service) GetByID(ctx context.Context, id string) (domain.Product, error) {
	// We call the repository for get the product by id
	product, err := s.repository.GetByID(ctx, id)

	// If we have an error log it and return it
	if err != nil {
		log.Println("[ProductsService][GetByID] error getting product by ID", err)
		return domain.Product{}, err
	}

	// We return the product
	return product, nil
}

// Update is a function that calls the repository for update a product by Id
func (s *service) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error) {
	// We call the repository for update the product by id
	product, err := s.repository.Update(ctx, product, id)

	// If we have an error log it and return it
	if err != nil {
		log.Println("[ProductsService][Update] error updating product by ID", err)
		return domain.Product{}, err
	}

	// We return the updated product
	return product, nil
}

// Delete is a function that calls the repository for delete a product by Id
func (s *service) Delete(ctx context.Context, id string) error {
	// We call the repository for delete the product by id
	err := s.repository.Delete(ctx, id)

	// If we have an error log it and return it
	if err != nil {
		log.Println("[ProductsService][Delete] error deleting product by ID", err)
		return err
	}

	// We return nill because we didn't have an error
	return nil
}
