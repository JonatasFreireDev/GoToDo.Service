package gorm

import (
	"github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
	"github.com/JonatasFreireDev/GoToDo.Service/src/shared/infra"
)

type TodoRepoType struct {
	todo model.Todo
}

func (r *TodoRepoType) Get(id int) (model.Todo, error) {
	result := infra.DB.First(&r.todo, id)

	if result.Error != nil {
		return model.Todo{}, result.Error
	}

	return r.todo, nil
}

func (r *TodoRepoType) GetAll() ([]model.Todo, error) {
	//TODO
}

func (r *TodoRepoType) Add(todo model.Todo) error {
	//TODO
}

func (r *TodoRepoType) Update(id int, todo model.Todo) error {
	//TODO
}

func (r *TodoRepoType) Delete(id int) error {
	//TODO
}

//Use Example
//var TodoRepository repository.TodoRepo = &TodoRepoType{}
