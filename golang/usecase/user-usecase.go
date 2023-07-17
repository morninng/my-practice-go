package usecase

import (
	"go-rest-api/model"
	"go-rest-api/models"
	"go-rest-api/repository"
	"go-rest-api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/volatiletech/null/v8"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	SignUp(user models.User) (model.UserResponse, error)
	Login(user models.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) SignUp(user models.User) (model.UserResponse, error) {

	if err := uu.uv.UserValidate(user); err != nil {
		return model.UserResponse{}, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password.String), 10)
	if err != nil {
		return model.UserResponse{}, err
	}
	newUser := models.User{Email: user.Email, Password: null.StringFrom(string(hash))}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    newUser.ID,
		Email: newUser.Email.String,
	}
	return resUser, nil
}

func (uu *userUsecase) Login(user models.User) (string, error) {

	storedUser := models.User{}
	if err := uu.ur.GetUserByEmail(&storedUser, user.Email.String); err != nil {
		return "", err
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password.String), []byte(user.Password.String))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": storedUser.ID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
