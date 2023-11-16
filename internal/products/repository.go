package products

import (
	"context"
	"errors"

	"github.com/burgosfacundo/ApiGo.git/internal/domain"
)

// Errors that can be returned in the response
var (
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("product not found")
)

// Repository represents a contract with all the functions that need to be implemented
type Repository interface {
	Create(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id string) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id string) (domain.Product, error)
	Delete(ctx context.Context, id string) error
}

// repository is a struct that contains the db of Product
type repository struct {
	db []domain.Product
}

// NewMemoryRepository is a function that loads the db into the repository
// because we still don't have a db sql connection
func NewMemoryRepository(db []domain.Product) Repository {
	return &repository{db: db}
}

// Create is a function that creates a new Product in the db
func (r *repository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	r.db = append(r.db, product)
	return product, nil
}

// GetAll is a function that returns all the products in the db
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	if len(r.db) < 1 {
		return []domain.Product{}, ErrEmpty
	}

	return r.db, nil
}

// GetByID is a function that returns a Product by id from the db
func (r *repository) GetByID(ctx context.Context, id string) (domain.Product, error) {
	var result domain.Product
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}

	if result.Id == "" {
		return domain.Product{}, ErrNotFound
	}

	return result, nil
}

// Update is a function that updates a Product by id from the db
func (r *repository) Update(
	ctx context.Context,
	product domain.Product,
	id string) (domain.Product, error) {

	var result domain.Product
	for key, value := range r.db {
		if value.Id == id {
			product.Id = id
			r.db[key] = product
			result = r.db[key]
			break
		}
	}

	if result.Id == "" {
		return domain.Product{}, ErrNotFound
	}

	return result, nil

}

// Delete is a function that deletes a Product by id from the db
func (r *repository) Delete(ctx context.Context, id string) error {
	var result domain.Product
	for key, value := range r.db {
		if value.Id == id {
			result = r.db[key]
			r.db = append(r.db[:key], r.db[key+1:]...)
			break
		}
	}

	if result.Id == "" {
		return ErrNotFound
	}

	return nil
}
