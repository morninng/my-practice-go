package controller

import (
	"encoding/json"
	"fmt"
	"go-rest-api/mock"
	"go-rest-api/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	taskJSON     = `{"title":"aaa"}`
	taskResponse = model.TaskResponse{ID: 1, Title: "aaa"}
)

func TestCreateTask(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockTaskUsecase := mock.NewMockITaskUsecase(mockCtrl)
	// テスト中に呼び出されるメソッドの振る舞いを定義
	mockTaskUsecase.
		EXPECT().
		CreateTask(gomock.Any()).
		Return(taskResponse, nil)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(taskJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := e.NewContext(req, rec)
	taskController := NewTaskController(mockTaskUsecase)

	// Assertions
	if assert.NoError(t, taskController.CreateTask(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		recBody := rec.Body
		fmt.Printf("%T\n", recBody)

		// *bytes.Buffer型のrecBodyをmodel.TaskResponse型に変換
		resultResponse := model.TaskResponse{}
		json.NewDecoder(recBody).Decode(&resultResponse)

		assert.Equal(t, taskResponse.Title, resultResponse.Title)
		assert.Equal(t, taskResponse.ID, resultResponse.ID)
	}
}
