package app

import (
	"web-endterm/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	app     *fiber.App
	Service service.Service
}

func NewServer(service service.Service) *Server{
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024,
	})

	app.Use(logger.New())

	s := &Server{
		app: app,
		Service: service,
	}

	s.Route()

	return s
}

func (s *Server) Start(port string) error {
	return s.app.Listen(":" + port)
}