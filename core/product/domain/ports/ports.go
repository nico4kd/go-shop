package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/product/domain/models"
)

type ProductRepository interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
}

type ProductApplication interface {
	Create(product *models.Product) error
	GetByID(id string) (*models.Product, error)
	GetAll() ([]*models.Product, error)
	Update(product *models.Product) error
	Delete(id string) error
}

type ProductHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
