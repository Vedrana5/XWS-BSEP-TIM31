package service

import (
	"github.com/google/uuid"
	"user/module/model"
	"user/module/repository"
)

type ConfirmationTokenService struct {
	Repo *repository.ConfirmationTokenRepository
}

func (service *ConfirmationTokenService) CreateConfirmationToken(confirmationToken *model.ConfirmationToken) error {
	err := service.Repo.CreateConfirmationToken(confirmationToken)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConfirmationTokenService) FindByToken(token uuid.UUID) *model.ConfirmationToken {
	confirmationToken := service.Repo.FindByToken(token)
	return confirmationToken
}

func (service *ConfirmationTokenService) UpdateConfirmationTokenValidity(token uuid.UUID, isValid bool) error {
	err := service.Repo.UpdateConfirmationTokenValidity(token, isValid)
	if err != nil {
		return err
	}
	return nil
}

func (service *ConfirmationTokenService) CheckIfHasAccess(token string) (interface{}, error) {

}
