package infra

import (
	todoModel "github.com/JonatasFreireDev/GoToDo.Service/src/modules/todo/model"
	userModel "github.com/JonatasFreireDev/GoToDo.Service/src/modules/user/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDatabase() {
	DB, err = gorm.Open(sqlite.Open("gotodo.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&todoModel.Todo{})
	DB.AutoMigrate(&userModel.User{})
}
