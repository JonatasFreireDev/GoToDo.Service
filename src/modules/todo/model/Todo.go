package model

import "time"

type Todo struct {
	Id          string    `json:"id" validate:"required"`
	Status      Status    `json:"status"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

var TodoList []Todo
