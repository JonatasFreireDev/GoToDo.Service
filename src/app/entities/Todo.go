package entities

import "time"

type Todo struct {
	Id          string    `json:"id"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var TodoList []Todo
