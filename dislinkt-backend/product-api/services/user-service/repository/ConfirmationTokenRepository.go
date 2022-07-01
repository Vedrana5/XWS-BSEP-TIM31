package repository

import (
	"fmt"
	"user/module/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ConfirmationTokenRepository struct {
	Database *gorm.DB
}

func (repo *ConfirmationTokenRepository) CreateConfirmationToken(confirmationToken *model.ConfirmationToken) error {
	result := repo.Database.Create(confirmationToken)
	fmt.Print(result)
	return nil
}

func (repo *ConfirmationTokenRepository) FindByToken(token uuid.UUID) *model.ConfirmationToken {
	confirmationToken := &model.ConfirmationToken{}
	repo.Database.First(&confirmationToken, "confirmation_token = ?", token)
	return confirmationToken
}

func (repo *ConfirmationTokenRepository) UpdateConfirmationTokenValidity(token uuid.UUID, isValid bool) error {
	result := repo.Database.Model(&model.ConfirmationToken{}).Where("confirmation_token = ?", token).Update("is_valid", isValid)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
