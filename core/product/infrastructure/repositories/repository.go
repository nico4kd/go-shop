package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-shop/core/product/domain/models"
	"github.com/solrac97gr/go-shop/core/product/domain/ports"
)

type ProductRepository struct {
	Database *sql.DB
}

var _ ports.ProductRepository = &ProductRepository{}

func NewProductRepository(database *sql.DB) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (p ProductRepository) Create(product *models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetByID(id string) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) GetAll() ([]*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) Update(product *models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
