package main

import "github.com/JonatasFreireDev/GoToDo.Service/src/shared/infra"

func main() {
	infra.App.Listen(":3000")
}
