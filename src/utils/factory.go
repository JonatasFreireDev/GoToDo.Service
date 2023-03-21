package utils

import (
	"time"

	"github.com/JonatasFreireDev/GoToDo.Service/src/app/entities"
	"github.com/gofiber/fiber/v2/utils"
)

func TodoFactory() entities.Todo {
	return entities.Todo{
		Id:          utils.UUIDv4(),
		Status:      entities.Status{Value: "pending"},
		Description: "description",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
