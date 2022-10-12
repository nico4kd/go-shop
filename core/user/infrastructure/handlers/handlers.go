package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/user/domain/ports"
)

type UserHandlers struct {
	application ports.UserApplication
}

var _ ports.UserHandler = (*UserHandlers)(nil)

func NewUserHandlers(application ports.UserApplication) *UserHandlers {
	return &UserHandlers{application: application}
}

func (u UserHandlers) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserHandlers) GetByID(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserHandlers) GetByEmail(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserHandlers) GetAll(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserHandlers) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (u UserHandlers) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
