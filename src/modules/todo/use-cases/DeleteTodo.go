package usecases

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
)

func remove[T any](slice []T, s int) []T {
	return append(slice[:s], slice[s+1:]...)
}

func DeleteTodo(id string) {
	for i, todo := range model.TodoList {
		if id == todo.Id {
			model.TodoList = remove(model.TodoList, i)
		}
	}
}
