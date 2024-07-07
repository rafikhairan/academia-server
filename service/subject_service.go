package service

import (
	"github.com/rafikhairan/academia/exception"
	"github.com/rafikhairan/academia/helper"
	"github.com/rafikhairan/academia/model"
	"gorm.io/gorm"
)

type SubjectService struct {
	DB *gorm.DB
}

func NewSubjectService(db *gorm.DB) *SubjectService {
	return &SubjectService{
		DB: db,
	}
}

func toSubjectResponse(subject model.Subject) model.SubjectResponse {
	return model.SubjectResponse{
		ID:          subject.ID,
		Name:        subject.Name,
		Description: subject.Description,
	}
}

func checkSubjectMustExists(db *gorm.DB, id string) model.Subject {
	var subject model.Subject

	if result := db.Take(&subject, "id = ?", id); result.RowsAffected == 0 {
		panic(exception.NewNotFoundError("Specialization not found"))
	}

	return subject
}

func (service *SubjectService) Create(request model.CreateSubjectRequest) model.SubjectResponse {
	subject := model.Subject{
		Name:        request.Name,
		Description: request.Description,
	}
	result := service.DB.Create(&subject)
	helper.PanicIfError(result.Error)

	return toSubjectResponse(subject)
}

func (service *SubjectService) GetAll() []model.SubjectResponse {
	var subjects []model.SubjectResponse

	result := service.DB.Model(&model.Subject{}).Find(&subjects)
	helper.PanicIfError(result.Error)

	return subjects
}

func (service *SubjectService) GetOne(id string) model.SubjectResponse {
	subject := checkSubjectMustExists(service.DB, id)

	return toSubjectResponse(subject)
}

func (service *SubjectService) Update(request model.UpdateSubjectRequest) model.SubjectResponse {
	subject := checkSubjectMustExists(service.DB, request.ID)

	subject.Name = request.Name
	subject.Description = request.Description

	result := service.DB.Save(&subject)
	helper.PanicIfError(result.Error)

	return toSubjectResponse(subject)
}

func (service *SubjectService) Delete(id string) {
	subject := checkSubjectMustExists(service.DB, id)

	result := service.DB.Delete(&subject)
	helper.PanicIfError(result.Error)
}
