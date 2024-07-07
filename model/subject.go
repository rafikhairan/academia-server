package model

import "github.com/google/uuid"

type Subject struct {
	BaseModel
	Name            string
	Description     *string
	Specializations []Specialization `gorm:"many2many:subject_specializations"`
}

type SubjectResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type CreateSubjectRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}

type UpdateSubjectRequest struct {
	ID          string  `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description"`
}
