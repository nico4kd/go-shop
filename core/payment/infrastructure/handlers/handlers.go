package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/core/payment/domain/ports"
)

type PaymentHandler struct {
	application ports.PaymentApplication
}

var _ ports.PaymentHandler = &PaymentHandler{}

func NewPaymentHandlers(application ports.PaymentApplication) *PaymentHandler {
	return &PaymentHandler{application: application}
}

func (p PaymentHandler) Create(cxt *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentHandler) GetByID(cxt *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentHandler) GetAll(cxt *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentHandler) Update(cxt *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentHandler) Delete(cxt *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
