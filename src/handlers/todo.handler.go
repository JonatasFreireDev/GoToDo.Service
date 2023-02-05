package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetTodo(c *fiber.Ctx) error {
	fmt.Printf("aew handle todo")
	return c.SendString("Hello, World!")
}
