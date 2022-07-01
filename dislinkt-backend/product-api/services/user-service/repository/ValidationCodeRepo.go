package repository

import (
	"fmt"
	"user/module/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ValidationCodeRepo struct {
	Database *gorm.DB
}

func (repo *ValidationCodeRepo) CreateConfirmationToken(validationCode *model.ValidationCode) error {
	result := repo.Database.Create(validationCode)
	fmt.Print(result)
	return nil
}

func (repo *ValidationCodeRepo) FindByCode(code uuid.UUID) *model.ValidationCode {
	validationCode := &model.ValidationCode{}
	repo.Database.First(&validationCode, "code = ?", code)
	return validationCode
}

func (repo *ValidationCodeRepo) UpdateValidationCodeValidity(token uuid.UUID, isValid bool) error {
	result := repo.Database.Model(&model.ValidationCode{}).Where("code = ?", token).Update("is_valid", isValid)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *ValidationCodeRepo) UpdateValidationCodeUsing(token uuid.UUID) error {
	result := repo.Database.Model(&model.ValidationCode{}).Where("code = ?", token).Update("is_used", true)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
