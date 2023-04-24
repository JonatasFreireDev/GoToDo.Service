package usecases

import "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"

func DeleteAllTodos() {
	model.TodoList = nil
}
