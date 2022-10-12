package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/user/domain/models"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type UserApplication interface {
	Create(user *models.User) error
	GetByID(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}

type UserHandler interface {
	Create(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	GetByEmail(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
