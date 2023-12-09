package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	// DB接続
	db := db.NewDB()
	// ユーザ
	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository, userValidator)
	userController := controller.NewUserController(userUseCase)
	// タスク
	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUseCase)
	// ルーティング
	e := router.NewRouter(userController, taskController)
	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
