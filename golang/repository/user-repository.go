package repository

import (
	"context"
	"database/sql"
	"go-rest-api/models"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserRepository interface {
	GetUserByEmail(user *models.User, email string) error
	CreateUser(user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(user *models.User, email string) error {

	sss := null.StringFrom(email)
	ctx := context.Background()
	var err error
	foundUser, err := models.Users(models.UserWhere.Email.EQ(sss)).One(ctx, ur.db)
	user.ID = foundUser.ID
	user.Email = foundUser.Email
	user.Password = foundUser.Password
	user.CreatedAt = foundUser.CreatedAt
	user.UpdatedAt = foundUser.UpdatedAt

	return err
}

func (ur *userRepository) CreateUser(user *models.User) error {
	ctx := context.Background()
	return user.Insert(ctx, ur.db, boil.Infer())

}
