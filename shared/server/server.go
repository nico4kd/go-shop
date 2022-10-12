package server

import (
	"github.com/gofiber/fiber/v2"
	orderHdl "github.com/solrac97gr/go-shop/core/order/domain/ports"
	paymentHdl "github.com/solrac97gr/go-shop/core/payment/domain/ports"
	productHdl "github.com/solrac97gr/go-shop/core/product/domain/ports"
	userHdl "github.com/solrac97gr/go-shop/core/user/domain/ports"
	"github.com/solrac97gr/go-shop/shared/middleware"

	"log"
)

type Server interface {
	Start(port string)
}

type FiberServer struct {
	UserHandlers    userHdl.UserHandler
	ProductHandlers productHdl.ProductHandler
	PaymentHandlers paymentHdl.PaymentHandler
	OrderHandlers   orderHdl.OrderHandler
	Middlewares     middleware.Middleware
}

var _ Server = (*FiberServer)(nil)

func NewFiberServer(
	userHandlers userHdl.UserHandler,
	productHandlers productHdl.ProductHandler,
	paymentHandlers paymentHdl.PaymentHandler,
	orderHandlers orderHdl.OrderHandler,
	middlewares middleware.Middleware,

) *FiberServer {
	return &FiberServer{
		UserHandlers:    userHandlers,
		ProductHandlers: productHandlers,
		PaymentHandlers: paymentHandlers,
		OrderHandlers:   orderHandlers,
		Middlewares:     middlewares,
	}
}

func (s *FiberServer) Start(port string) {
	app := fiber.New()

	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Get("/:id", s.UserHandlers.GetByID)
	userRoutes.Get("/", s.UserHandlers.GetAll)
	userRoutes.Post("/", s.UserHandlers.Create)
	userRoutes.Put("/:id", s.UserHandlers.Update)
	userRoutes.Delete("/:id", s.UserHandlers.Delete)

	productRoutes := v1.Group("/product")
	productRoutes.Get("/:id", s.ProductHandlers.GetByID)
	productRoutes.Get("/", s.ProductHandlers.GetAll)
	productRoutes.Post("/", s.ProductHandlers.Create)
	productRoutes.Put("/:id", s.ProductHandlers.Update)
	productRoutes.Delete("/:id", s.ProductHandlers.Delete)

	paymentRoutes := v1.Group("/payment")
	paymentRoutes.Get("/:id", s.PaymentHandlers.GetByID)
	paymentRoutes.Get("/", s.PaymentHandlers.GetAll)
	paymentRoutes.Post("/", s.PaymentHandlers.Create)
	paymentRoutes.Put("/:id", s.PaymentHandlers.Update)
	paymentRoutes.Delete("/:id", s.PaymentHandlers.Delete)

	orderRoutes := v1.Group("/order")
	orderRoutes.Get("/:id", s.OrderHandlers.GetByID)
	orderRoutes.Get("/", s.OrderHandlers.GetAll)
	orderRoutes.Post("/", s.OrderHandlers.Create)
	orderRoutes.Put("/:id", s.OrderHandlers.Update)
	orderRoutes.Delete("/:id", s.OrderHandlers.Delete)

	if port == "" {
		port = "3000"
	}

	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
