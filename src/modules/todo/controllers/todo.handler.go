package controllers

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
	usecases "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/use-cases"
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/utils"
	"github.com/gofiber/fiber/v2"
)

func GetTodo(c *fiber.Ctx) error {
	return c.JSON(model.TodoList)
}

func PostTodo(c *fiber.Ctx) error {
	todo := utils.TodoFactory()

	if err := c.BodyParser(&todo); err != nil {
		return err
	}

	usecases.AddNewTodo(todo)

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")

	usecases.DeleteTodo(id)

	return c.SendStatus(fiber.StatusOK)
}
