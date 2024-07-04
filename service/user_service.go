package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rafikhairan/academia/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (service *UserService) Register(request model.RegisterRequest) (model.User, error) {
	var user model.User

	result := service.db.Where("email = ?", request.Email).First(&user)
	if result.RowsAffected >= 1 {
		return user, fiber.NewError(400, "register failed")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return user, err
	}

	user = model.User{
		Email:    request.Email,
		Password: string(hashPassword),
	}

	if result = service.db.Create(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (service *UserService) Login(request model.LoginRequest) (model.User, error) {
	var user model.User

	result := service.db.Select("id", "email", "password").Where("email = ?", request.Email).First(&user)
	if result.RowsAffected == 0 {
		return user, fiber.NewError(404, "email not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return user, fiber.NewError(401, "password incorrect")
	}

	return user, nil
}
