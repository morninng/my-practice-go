package usecase

import (
	"go-rest-api/model"
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITaskUsecase interface {
	GetAllTasks() ([]model.TaskResponse, error)
	GetTaskById(taskId int64) (model.TaskResponse, error)
	CreateTask(task models.Task) (model.TaskResponse, error)
	UpdateTask(task models.Task, taskId int64) (model.TaskResponse, error)
	DeleteTask(taskId int64) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks() ([]model.TaskResponse, error) {
	// tasks := []model.Task{}
	tasks, err := tu.tr.GetAllTasks()
	if err != nil {
		return nil, err
	}
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

func (tu *taskUsecase) GetTaskById(taskId int64) (model.TaskResponse, error) {

	task, err := tu.tr.GetTaskById(taskId)
	if err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task models.Task) (model.TaskResponse, error) {

	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task models.Task, taskId int64) (model.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return model.TaskResponse{}, err
	}
	if err := tu.tr.UpdateTask(&task, taskId); err != nil {
		return model.TaskResponse{}, err
	}
	resTask := model.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(taskId int64) error {
	if err := tu.tr.DeleteTask(taskId); err != nil {
		return err
	}
	return nil
}
