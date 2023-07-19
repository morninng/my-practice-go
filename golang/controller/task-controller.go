package controller

import (
	"fmt"
	"go-rest-api/models"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	fmt.Println("--------------------")

	tasksRes, err := tc.tu.GetAllTasks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(c echo.Context) error {

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(int64(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(c echo.Context) error {
	fmt.Println("CreateTask")

	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println("task", task)
	taskRes, err := tc.tu.CreateTask(task)
	fmt.Println("taskResponse", taskRes)
	fmt.Printf("taskResponse %T \n", taskRes)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	task := models.Task{}
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, int64(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {

	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(int64(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
