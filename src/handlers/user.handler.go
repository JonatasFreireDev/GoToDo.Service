package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	fmt.Printf("aew handle user")
	return c.SendString("Get user!")
}
