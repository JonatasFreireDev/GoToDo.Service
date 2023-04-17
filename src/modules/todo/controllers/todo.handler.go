package controllers

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
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

	model.TodoList = append(model.TodoList, todo)

	return c.SendStatus(fiber.StatusCreated)
}
