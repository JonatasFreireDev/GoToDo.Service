package usecases

import "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"

func AddNewTodo(todo model.Todo) {
	model.TodoList = append(model.TodoList, todo)
}
