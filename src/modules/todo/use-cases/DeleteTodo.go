package usecases

import (
	"errors"

	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
)

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteTodo(id string) error {
	for i, todo := range model.TodoList {
		if id == todo.Id {
			model.TodoList = remove(model.TodoList, i)
			return nil
		}
	}

	return errors.New("nao foi possivel deletar")
}
