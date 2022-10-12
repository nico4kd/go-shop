package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/order/domain/ports"
)

type OrderHandler struct {
	OrderApplication ports.OrderApplication
}

var _ ports.OrderHandler = &OrderHandler{}

func NewOrderHandlers(orderApplication ports.OrderApplication) ports.OrderHandler {
	return &OrderHandler{OrderApplication: orderApplication}
}

func (o OrderHandler) Create(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderHandler) GetByID(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderHandler) GetAll(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderHandler) Update(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderHandler) Delete(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
