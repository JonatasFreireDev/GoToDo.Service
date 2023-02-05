package main

import "github.com/JonatasFreireDev/GoToDo.Service/src/routes"

func main() {
	routes.App.Listen(":3000")
}
