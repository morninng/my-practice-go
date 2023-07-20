package main

import (
	"fmt"
	"go-rest-api/router"
)

func main() {
	fmt.Println("my practice go")

	// db := db.NewDB()
	// taskValidator := validator.NewTaskValidator()
	// userValidator := validator.NewUserValidator()
	// userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	// userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	// userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)

	taskController := InitializeTaskController()
	userController := InitializeUserController()

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))

}
