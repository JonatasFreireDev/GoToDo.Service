package routes

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r *fiber.App) {
	api := r.Group("/todo")

	api.Get("/", controllers.GetTodo)
	api.Post("/", controllers.PostTodo)
}
