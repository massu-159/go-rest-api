package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskByID(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

// 構造体を作成
type taskController struct {
	tu usecase.ITaskUseCase
}

// コンストラクタを作成
// 依存関係の注入を行う(di)
// 構造体の実体を作成する
func NewTaskController(tu usecase.ITaskUseCase) ITaskController {
	return &taskController{tu}
}

// タスクを全て取得する
func (tc *taskController) GetAllTasks(c echo.Context) error {
	// ユーザーIDを取得する
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	// タスクを取得する
	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

// タスクを取得する
func (tc *taskController) GetTaskByID(c echo.Context) error {
	// ユーザーIDを取得する
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	// タスクIDを取得する
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスクを取得する
	taskRes, err := tc.tu.GetTaskByID(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを作成する
func (tc *taskController) CreateTask(c echo.Context) error {
	// リクエストボディを取得する
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザーIDを取得する
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	task.UserId = uint(userId.(float64))
	// タスクを作成する
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを更新する
func (tc *taskController) UpdateTask(c echo.Context) error {
	// リクエストボディを取得する
	task := model.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// ユーザーIDを取得する
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	// タスクIDを取得する
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスクを更新する
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクを削除する
func (tc *taskController) DeleteTask(c echo.Context) error {
	// ユーザーIDを取得する
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["userId"]

	// タスクIDを取得する
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	// タスクを削除する
	if err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
