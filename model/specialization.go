package model

import "github.com/google/uuid"

type Specialization struct {
	BaseModel
	Name        string
	Description string
}

type SpecializationData struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type SpecializationRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}
