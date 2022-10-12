package application

import (
	"github.com/solrac97gr/go-shop/core/product/domain/models"
	"github.com/solrac97gr/go-shop/core/product/domain/ports"
)

type ProductApplication struct {
	ProductRepository ports.ProductRepository
}

var _ ports.ProductApplication = &ProductApplication{}

func NewProductApplication(productRepository ports.ProductRepository) *ProductApplication {
	return &ProductApplication{
		ProductRepository: productRepository,
	}
}

func (p ProductApplication) Create(product *models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductApplication) GetByID(id string) (*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductApplication) GetAll() ([]*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductApplication) Update(product *models.Product) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductApplication) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
