package utils

import (
	"time"

	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
	"github.com/gofiber/fiber/v2/utils"
)

func TodoFactory() model.Todo {
	return model.Todo{
		Id:          utils.UUIDv4(),
		Status:      model.Status{Value: "pending", UpdatedAt: time.Now()},
		Description: "description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
