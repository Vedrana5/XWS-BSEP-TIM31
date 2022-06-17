package service

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/repository"
	"github.com/google/uuid"
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
