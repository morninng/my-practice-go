package repository

import (
	"database/sql"
	"go-rest-api/model"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 普通に、データベースを利用してテストをしてみる
// gomockを利用するのに、生のsqlを調べる必要があるので。
/*
func TestGetAllTasks(t *testing.T) {
	db := db.NewDB()
	taskRepository := NewTaskRepository(db)
	tasks, _ := taskRepository.GetAllTasks()
	fmt.Println("----------", tasks)

}
*/

// mockを利用してテストをしてみる

// gorm の mockDBの作成 https://zenn.dev/ymktmk/articles/27668df082a9b2
func NewDbMock() (*sql.DB, *gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	mockDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return sqlDB, mockDB, mock, err
}

func TestGetAllTasksFromRepository(t *testing.T) {

	allTasks := []model.Task{{ID: 1, Title: "test1", CreatedAt: time.Now(), UpdatedAt: time.Now()}, {ID: 2, Title: "test2", CreatedAt: time.Now(), UpdatedAt: time.Now()}}

	sqlDB, db, mock, err := NewDbMock()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "tasks" ORDER BY created_at`)).
		WillReturnRows(mock.NewRows([]string{"ID", "Title", "CreatedAt", "UpdatedAt"}).
			AddRow(allTasks[0].ID, allTasks[0].Title, allTasks[0].CreatedAt, allTasks[0].UpdatedAt).
			AddRow(allTasks[1].ID, allTasks[1].Title, allTasks[1].CreatedAt, allTasks[1].UpdatedAt))

	mock.ExpectCommit()

	taskRepository := NewTaskRepository(db)
	resultTasks, _ := taskRepository.GetAllTasks()

	if !reflect.DeepEqual(resultTasks, allTasks) {
		t.Error("tasks are not equal")
	}

}
