package repository

import "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"

type TodoRepo interface {
	Get(id int) (model.Todo, error)
	GetAll() ([]model.Todo, error)
	Add(todo model.Todo) error
	Update(id int, todo model.Todo) error
	Delete(id int) error
}
