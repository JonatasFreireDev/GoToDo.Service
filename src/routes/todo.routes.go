package routes

import "github.com/JonatasFreireDev/GoToDo.Service/src/handlers"

func init() {
	App.Get("/", handlers.GetTodo)
}
