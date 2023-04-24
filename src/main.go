package main

import "github.com/JonatasFreireDev/GoToDo.Service/src/shared/infra"

func main() {
	infra.InitRoutes()
	infra.InitDatabase()

	infra.App.Listen(":3000")
}
