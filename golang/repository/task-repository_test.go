package repository

import (
	"fmt"
	"go-rest-api/models"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/boil"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

// 普通に、データベースを利用してテストをしてみる
// gomockを利用するのに、生のsqlを調べる必要があるので。
/*
func TestGetAllTasks(t *testing.T) {
	db := db.NewDB()
	taskRepository := NewTaskRepository(db)
	tasks, _ := taskRepository.GetAllTasks()
	fmt.Println("---- tasks ------", tasks)

}*/

// mockを利用してテストをしてみる

// gorm の mockDBの作成 https://zenn.dev/ymktmk/articles/27668df082a9b2
// func NewDbMock() (*sql.DB, *sql.DB, sqlmock.Sqlmock, error) {
// 	sqlDB, mock, err := sqlmock.New()
// 	mockDB, err := sql.Open(postgres.New(postgres.Config{
// 		Conn: sqlDB,
// 	}), &gorm.Config{})
// 	return sqlDB, mockDB, mock, err
// }

func TestGetAllTasksFromRepository(t *testing.T) {

	now := null.TimeFrom(time.Now())

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	boil.SetDB(db)

	allTasks := []models.Task{{ID: 1, Title: "abc1-abc1-abc1", CreatedAt: now, UpdatedAt: now}, {ID: 2, Title: "abc2-abc2-abc2", CreatedAt: now, UpdatedAt: now}}

	rows := sqlmock.NewRows([]string{"id", "title", "created_at", "updated_at"}).
		AddRow(allTasks[0].ID, allTasks[0].Title, allTasks[0].CreatedAt, allTasks[0].UpdatedAt).
		AddRow(allTasks[1].ID, allTasks[1].Title, allTasks[1].CreatedAt, allTasks[1].UpdatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT "tasks".* FROM "tasks"`)).
		WillReturnRows(rows)

	taskRepository := NewTaskRepository(db)
	resultTasks, err := taskRepository.GetAllTasks()
	if err != nil {
		fmt.Println("error 11 ", err)
	}
	fmt.Println("resultTasks", resultTasks)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf(">>> Unfulfilled expectations: %s", err)
	}
	fmt.Println("<<<<< mock call expectation matched")

	if !reflect.DeepEqual(resultTasks, allTasks) {
		t.Error(">>> tasks are not equal")
	}
	fmt.Println("tasks are equqal")

}
