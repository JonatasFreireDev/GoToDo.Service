package routes

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/handlers"
)

func todoRoutes() {
	api := App.Group("/todo")

	api.Get("/", handlers.GetTodo)
}
