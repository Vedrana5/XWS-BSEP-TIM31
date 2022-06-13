package dto

import "github.com/google/uuid"

type EditProfileDTO struct {
	ID             uuid.UUID `json:"ID" validate:"required"`
	OldUsername    string    `json:"OldUsername" validate:"required"`
	Username       string    `json:"Username" validate:"required"`
	Password       string    `json:"Password" validate:"required"`
	Email          string    `json:"Email" validate:"required,email"`
	PhoneNumber    string    `json:"PhoneNumber" validate:"required"`
	FirstName      string    `json:"FirstName" validate:"required"`
	LastName       string    `json:"LastName" validate:"required"`
	DateOfBirth    string    `json:"DateOfBirth" validate:"required"`
	TypeOfUser     string    `json:"User_type" gorm:"not null"`
	TypeOfProfile  string    `json:"Profile_type" gorm:"not null"`
	Gender         string    `json:"Gender" gorm:"not null"`
	Biography      string    `json:"Biography" gorm:"not null"`
	WorkExperience string    `json:"WorkExperience" gorm:"not null"`
	Education      string    `json:"Education" gorm:"not null"`
	Skills         string    `json:"Skills" gorm:"not null"`
	Interest       string    `json:"Interest" gorm:"not null"`
	Question       string    `json:"Question"`
	Answer         string    `json:"Answer"`
}
