package infra

import (
	todo "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/routes"
	"github.com/gofiber/fiber/v2"
)

var App *fiber.App

func init() {
	App = fiber.New()

	todo.RegisterRoutes(App)
}
