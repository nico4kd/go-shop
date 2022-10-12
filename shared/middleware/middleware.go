package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/solrac97gr/go-shop/shared/auth"
)

type Middleware interface {
}

type FiberMiddleware struct {
	authenticator auth.Authenticator
}

func NewFiberMiddleware(auth auth.Authenticator) *FiberMiddleware {
	return &FiberMiddleware{
		authenticator: auth,
	}
}

func (m *FiberMiddleware) PrivateRoute(ctx *fiber.Ctx) error {
	credentials := new(auth.Credentials)

	if err := ctx.BodyParser(credentials); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err := m.authenticator.Authenticate(credentials)
	if err != nil {
		return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Next()
}
