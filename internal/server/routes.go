package server

import (
	"noid/internal/database"

	"github.com/gofiber/fiber/v2"
)

func (s *FiberServer) RegisterFiberRoutes() {
	s.App.Get("/", s.HelloWorldHandler)
	s.App.Get("/health", s.healthHandler)

	postR := s.App.Group("/api/v1").Post
	postR("/signup", s.signUpHandler)
}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

// sign up user
func (s *FiberServer) signUpHandler(c *fiber.Ctx) error {
	c.Accepts("application/json")
	u := new(database.User)

	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	err := s.db.CreateUser(u)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.JSON(u)
}
