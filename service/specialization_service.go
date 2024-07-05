package service

import (
	"github.com/gofiber/fiber/v2"
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
func (service *SpecializationService) Create(request model.SpecializationRequest) (model.SpecializationData, error) {
	var specializationData model.SpecializationData

	specialization := model.Specialization{
		Name:        request.Name,
		Description: request.Description,
	}

	if result := service.DB.Create(&specialization).Scan(&specializationData); result.Error != nil {
		return specializationData, result.Error
	}

	return specializationData, nil
}

func (service *SpecializationService) GetAll() ([]model.SpecializationData, error) {
	var specializations []model.SpecializationData

	if result := service.DB.Find(&[]model.Specialization{}).Scan(&specializations); result.Error != nil {
		return specializations, result.Error
	}

	return specializations, nil
}

func (service *SpecializationService) GetOne(id string) (model.SpecializationData, error) {
	var specializationData model.SpecializationData

	if result := service.DB.First(&model.Specialization{}, "id = ?", id).Scan(&specializationData); result.RowsAffected == 0 {
		return specializationData, fiber.NewError(404, "specialization not found")
	}

	return specializationData, nil
}

func (service *SpecializationService) Update(id string, request model.SpecializationRequest) (model.SpecializationData, error) {
	var specialization model.Specialization
	var specializationData model.SpecializationData

	if result := service.DB.First(&specialization, "id = ?", id); result.RowsAffected == 0 {
		return specializationData, fiber.NewError(404, "specialization not found")
	}

	specialization.Name = request.Name
	specialization.Description = request.Description

	if result := service.DB.Save(&specialization).Scan(&specializationData); result.Error != nil {
		return specializationData, result.Error
	}

	return specializationData, nil
}

func (service *SpecializationService) Delete(id string) (model.SpecializationData, error) {
	var specialization model.Specialization
	var specializationData model.SpecializationData

	if result := service.DB.First(&specialization, "id = ?", id).Scan(&specializationData); result.RowsAffected == 0 {
		return specializationData, fiber.NewError(404, "specialization not found")
	}

	service.DB.Delete(&specialization)

	return specializationData, nil
}
