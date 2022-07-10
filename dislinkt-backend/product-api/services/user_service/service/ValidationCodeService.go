package service

import (
	"github.com/google/uuid"
	"user/module/model"
	"user/module/repository"
)

type ValidationCodeService struct {
	Repo *repository.ValidationCodeRepo
}

func (service *ValidationCodeService) CreateConfirmationToken(validationCode *model.ValidationCode) error {
	err := service.Repo.CreateConfirmationToken(validationCode)
	if err != nil {
		return err
	}
	return nil
}

func (service *ValidationCodeService) FindByCode(token uuid.UUID) *model.ValidationCode {
	validationCode := service.Repo.FindByCode(token)
	return validationCode
}

func (service *ValidationCodeService) UpdateValidationCodeValidity(token uuid.UUID, isValid bool) error {
	err := service.Repo.UpdateValidationCodeValidity(token, isValid)
	if err != nil {
		return err
	}
	return nil
}

func (service *ValidationCodeService) UpdateValidationCodeUsing(token uuid.UUID) error {
	err := service.Repo.UpdateValidationCodeUsing(token)
	if err != nil {
		return err
	}
	return nil
}
