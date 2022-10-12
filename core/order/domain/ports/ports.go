package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/order/domain/models"
)

type OrderRepository interface {
	Create(order *models.Order) error
	GetByID(id int64) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(order *models.Order) error
}

type OrderApplication interface {
	Create(order *models.Order) error
	GetByID(id int64) (*models.Order, error)
	GetAll() ([]*models.Order, error)
	Update(order *models.Order) error
	Delete(order *models.Order) error
}

type OrderHandler interface {
	Create(ctx *fiber.Ctx) error
	GetByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
