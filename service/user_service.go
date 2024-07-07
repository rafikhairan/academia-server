package service

import (
	"github.com/rafikhairan/academia/auth"
	"github.com/rafikhairan/academia/exception"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (service *UserService) Register(request model.RegisterRequest) model.UserResponse {
	var user model.User

	result := service.DB.Take(&user, "email = ?", request.Email)
	if result.RowsAffected >= 1 {
		panic(exception.NewBadRequestError("Register failed"))
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	helper.PanicIfError(err)

	user = model.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	result = service.DB.Create(&user)
	helper.PanicIfError(result.Error)

	return model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

func (service *UserService) Login(request model.LoginRequest) model.UserResponse {
	var user model.User

	if result := service.DB.Where("email = ?", request.Email).Take(&user); result.RowsAffected == 0 {
		panic(exception.NewBadRequestError("Login failed"))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		panic(exception.NewBadRequestError("Login failed"))
	}

	return model.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Token: auth.GenerateToken(user),
	}
}
