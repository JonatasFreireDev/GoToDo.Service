package routes

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/user/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r *fiber.App) {
	api := r.Group("/user")

	api.Get("/", controllers.GetUser)
}
