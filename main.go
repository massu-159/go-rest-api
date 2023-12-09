package main

import (
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
)

func main() {
	// DB接続
	db := db.NewDB()
	// ユーザ
	userRepository := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepository)
	userController := controller.NewUserController(userUseCase)
	// タスク
	taskRepository := repository.NewTaskRepository(db)
	taskUseCase := usecase.NewTaskUseCase(taskRepository)
	taskController := controller.NewTaskController(taskUseCase)
	// ルーティング
	e := router.NewRouter(userController, taskController)
	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
