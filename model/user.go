package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return
}

type UserResponse struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

func (user *User) ToUserResponse() UserResponse {
	return UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}
}

type RegisterRequest struct {
	Email           string `validate:"required,email" json:"email"`
	Password        string `validate:"required,min=8" json:"password"`
	ConfirmPassword string `validate:"required,min=8,eqfield=Password" json:"confirm_password"`
}

type LoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8" json:"password"`
}
