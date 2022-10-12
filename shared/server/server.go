package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server interface {
	Start()
}

type FiberServer struct {
}

var _ Server = (*FiberServer)(nil)

func NewFiberServer() *FiberServer {
	return &FiberServer{}
}

func (s *FiberServer) Start() {
	app := fiber.New()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
