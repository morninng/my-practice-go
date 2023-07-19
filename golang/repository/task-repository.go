package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest-api/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ITaskRepository interface {
	GetAllTasks() ([]models.Task, error)
	GetTaskById(id int64) (*models.Task, error)
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task, id int64) error
	DeleteTask(id int64) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks() ([]models.Task, error) {

	ctx := context.Background()
	tasks, err := models.Tasks().All(ctx, tr.db)
	fmt.Println("!!!!tasks", tasks)
	if err != nil {
		return nil, err
	}

	aaa := []models.Task{}

	for _, v := range tasks {
		fmt.Println("v vvvvvvvvvvvvvvvvvvvvvv", *v)
		aaa = append(aaa, *v)
	}
	fmt.Println("!!!!aaa", aaa)
	return aaa, nil
}

func (tr *taskRepository) GetTaskById(id int64) (*models.Task, error) {
	ctx := context.Background()
	var err error
	task, err := models.Tasks(models.TaskWhere.ID.EQ(id)).One(ctx, tr.db)

	return task, err
}

func (tr *taskRepository) CreateTask(task *models.Task) error {
	ctx := context.Background()
	return task.Insert(ctx, tr.db, boil.Infer())

}

func (tr *taskRepository) UpdateTask(task *models.Task, id int64) error {

	ctx := context.Background()
	originalTask, err := models.FindTask(ctx, tr.db, id)
	if err != nil {
		return err
	}
	originalTask.Title = task.Title
	rowAff, err := originalTask.Update(ctx, tr.db, boil.Infer())
	if err != nil {
		return err
	}
	fmt.Println("rowAff", rowAff)
	return nil
}

func (tr *taskRepository) DeleteTask(id int64) error {

	ctx := context.Background()
	originalTask, err := models.FindTask(ctx, tr.db, id)
	if err != nil {
		return err
	}
	rowAff, err := originalTask.Delete(ctx, tr.db)
	if err != nil {
		return err
	}
	fmt.Println("rowAff", rowAff)
	return nil
}
