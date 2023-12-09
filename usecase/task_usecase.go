package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type ITaskUseCase interface {
	GetAllTasks(userId uint) ([]model.TaskResponse, error)
	GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error)
	CreateTask(task model.Task) (model.TaskResponse, error)
	UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

// 構造体を作成
type taskUseCase struct {
	tr repository.ITaskRepository
}

// コンストラクタを作成
// 依存関係の注入を行う(di)
// 構造体の実体を作成する
func NewTaskUseCase(tr repository.ITaskRepository) ITaskUseCase {
	return &taskUseCase{tr}
}

// タスクを全て取得する
func (tu *taskUseCase) GetAllTasks(userId uint) ([]model.TaskResponse, error) {
	// タスクを取得する
	tasks := []model.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	// レスポンス用のタスクに値を詰め替える
	resTasks := []model.TaskResponse{}
	for _, v := range tasks {
		t := model.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

// タスクを取得する
func (tu *taskUseCase) GetTaskByID(userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクを取得する
	task := model.Task{}
	if err := tu.tr.GetTaskByID(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	// レスポンス用のタスクに値を詰め替える
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

// タスクを作成する
func (tu *taskUseCase) CreateTask(task model.Task) (model.TaskResponse, error) {
	// タスクを作成する
	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	// レスポンス用のタスクに値を詰め替える
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

// タスクを更新する
func (tu *taskUseCase) UpdateTask(task model.Task, userId uint, taskId uint) (model.TaskResponse, error) {
	// タスクを更新する
	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	// レスポンス用のタスクに値を詰め替える
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

// タスクを削除する
func (tu *taskUseCase) DeleteTask(userId uint, taskId uint) error {
	// タスクを削除する
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
