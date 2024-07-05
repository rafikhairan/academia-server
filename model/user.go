package model

import "github.com/google/uuid"

type User struct {
	BaseModel
	Email    string
	Password string
}

type UserAuthenticationData struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type RegisterRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,eqfield=Password"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
