package controllers

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/app/entities"
	"github.com/JonatasFreireDev/GoToDo.Service/src/utils"
	"github.com/gofiber/fiber/v2"
)

func GetTodo(c *fiber.Ctx) error {
	return c.JSON(entities.TodoList)
}

func PostTodo(c *fiber.Ctx) error {
	todo := utils.TodoFactory()

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	entities.TodoList = append(entities.TodoList, todo)

	return c.SendStatus(fiber.StatusCreated)
}
