package server

import (
	"log"
	"world-api/routes"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	port   string
	server *fiber.App
}

func NewServer() Server {
	return Server{
		port:   "3333",
		server: fiber.New(),
	}
}

func (s *Server) Run() {
	routes.ConfigRoutes(s.server)

	if err := s.server.Listen(":3000"); err != nil {
		log.Printf("listen: %s\n", err)
	}
}
