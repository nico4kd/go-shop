package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/payment/domain/models"
)

type PaymentRepository interface {
	Create(payment *models.Payment) error
	GetByID(id string) (*models.Payment, error)
	GetAll() ([]*models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id string) error
}

type PaymentApplication interface {
	Create(payment *models.Payment) error
	GetByID(id string) (*models.Payment, error)
	GetAll() ([]*models.Payment, error)
	Update(payment *models.Payment) error
	Delete(id string) error
}

type PaymentHandler interface {
	Create(cxt *fiber.Ctx) error
	GetByID(cxt *fiber.Ctx) error
	GetAll(cxt *fiber.Ctx) error
	Update(cxt *fiber.Ctx) error
	Delete(cxt *fiber.Ctx) error
}
