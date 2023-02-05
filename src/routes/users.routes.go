package routes

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/handlers"
)

func userRoutes() {
	api := App.Group("/user")

	api.Get("/", handlers.GetUser)
}
