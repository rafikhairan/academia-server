package service

import (
	"github.com/rafikhairan/academia/exception"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"gorm.io/gorm"
)

type SpecializationService struct {
	DB *gorm.DB
}

func NewSpecializationService(db *gorm.DB) *SpecializationService {
	return &SpecializationService{
		DB: db,
	}
}

func toSpecializationResponse(specialization model.Specialization) model.SpecializationResponse {
	return model.SpecializationResponse{
		ID:          specialization.ID,
		Name:        specialization.Name,
		Description: specialization.Description,
	}
}

func checkSpecializationMustExists(db *gorm.DB, id string) model.Specialization {
	var specialization model.Specialization

	if result := db.Take(&specialization, "id = ?", id); result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("Specialization not found"))
	}

	return specialization
}

func (service *SpecializationService) Create(request model.CreateSpecializationRequest) model.SpecializationResponse {
	specialization := model.Specialization{
		Name:        request.Name,
		Description: request.Description,
	}
	result := service.DB.Create(&specialization)
	helper.PanicIfError(result.Error)

	return toSpecializationResponse(specialization)
}

func (service *SpecializationService) GetAll() []model.SpecializationResponse {
	var specializations []model.SpecializationResponse

	result := service.DB.Model(&model.Specialization{}).Find(&specializations)
	helper.PanicIfError(result.Error)

	return specializations
}

func (service *SpecializationService) GetOne(id string) model.SpecializationResponse {
	specialization := checkSpecializationMustExists(service.DB, id)

	return toSpecializationResponse(specialization)
}

func (service *SpecializationService) Update(request model.UpdateSpecializationRequest) model.SpecializationResponse {
	specialization := checkSpecializationMustExists(service.DB, request.ID)

	specialization.Name = request.Name
	specialization.Description = request.Description
	
	result := service.DB.Save(&specialization)
	helper.PanicIfError(result.Error)

	return toSpecializationResponse(specialization)
}

func (service *SpecializationService) Delete(id string) model.SpecializationResponse {
	specialization := checkSpecializationMustExists(service.DB, id)

	result := service.DB.Delete(&specialization)
	helper.PanicIfError(result.Error)

	return toSpecializationResponse(specialization)
}
