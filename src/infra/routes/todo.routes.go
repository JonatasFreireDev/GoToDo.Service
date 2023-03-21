package routes

import "github.com/JonatasFreireDev/GoToDo.Service/src/infra/controllers"

func todoRoutes() {
	api := App.Group("/todo")

	api.Get("/", controllers.GetTodo)
	api.Post("/", controllers.PostTodo)
}
