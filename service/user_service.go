package service

import (
	"github.com/gofiber/fiber/v2"
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

func (service *UserService) Register(request model.RegisterRequest) (model.UserAuthData, error) {
	var user model.User
	var userAuthData model.UserAuthData

	if result := service.DB.First(&user, "email = ?", request.Email); result.RowsAffected >= 1 {
		return userAuthData, fiber.NewError(400, "register failed")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return userAuthData, err
	}

	user = model.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	if result := service.DB.Create(&user).Scan(&userAuthData); result.Error != nil {
		return userAuthData, result.Error
	}

	return userAuthData, nil
}

func (service *UserService) Login(request model.LoginRequest) (model.UserAuthData, error) {
	var user model.User
	var userAuthData model.UserAuthData

	if result := service.DB.First(&user, "email = ?", request.Email).Scan(&userAuthData); result.RowsAffected == 0 {
		return userAuthData, fiber.NewError(404, "email not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return userAuthData, fiber.NewError(401, "password incorrect")
	}

	return userAuthData, nil
}
