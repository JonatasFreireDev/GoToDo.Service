package routes

import "github.com/JonatasFreireDev/GoToDo.Service/src/infra/controllers"

func userRoutes() {
	api := App.Group("/user")

	api.Get("/", controllers.GetUser)
}
