package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/product/domain/ports"
)

type ProductHandler struct {
	ProductApplication ports.ProductApplication
}

var _ ports.ProductHandler = &ProductHandler{}

func NewProductHandlers(productApplication ports.ProductApplication) *ProductHandler {
	return &ProductHandler{
		ProductApplication: productApplication,
	}
}

func (p ProductHandler) Create(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHandler) GetByID(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHandler) GetAll(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHandler) Update(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p ProductHandler) Delete(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
