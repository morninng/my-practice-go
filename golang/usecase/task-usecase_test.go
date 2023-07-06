package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository_mock"
	"go-rest-api/validator_mock"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestGetAllTasks(t *testing.T) {

	var (
		allTasks         = []model.Task{{ID: 1, Title: "title1"}, {ID: 2, Title: "title2"}}
		allTasksResponse = []model.TaskResponse{{ID: 1, Title: "title1"}, {ID: 2, Title: "title2"}}
	)

	// repository mock
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTaskRepsitory := repository_mock.NewMockITaskRepository(mockCtrl)

	mockTaskRepsitory.
		EXPECT().
		GetAllTasks().
		Return(allTasks, nil)

	// validator mock
	mockCtrl2 := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTaskValidator := validator_mock.NewMockITaskValidator(mockCtrl2)

	taskUsecase := NewTaskUsecase(mockTaskRepsitory, mockTaskValidator)

	resAllTasks, _ := taskUsecase.GetAllTasks()
	if !reflect.DeepEqual(resAllTasks, allTasksResponse) {
		t.Error("tasks are not equal")
	}

	// Assertions

}
