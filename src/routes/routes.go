package routes

import (
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func init() {
	App = fiber.New()
}

func Register() {
	todoRoutes()
	userRoutes()
}
