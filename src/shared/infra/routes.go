package infra

import (
	"log"

	todo "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/routes"
	"github.com/JonatasFreireDev/GoToDo.Service/src/shared/config"
	"github.com/gofiber/fiber/v2"
)

type routerToRegister func(*fiber.App)

var App *fiber.App

func InitRoutes() {
	App = fiber.New()

	register(todo.RegisterRoutes)
}

func register(f routerToRegister) {
	file, err := config.GetConf("..\\config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	if file.Config.RequestType == "http" {
		f(App)
	}
}
