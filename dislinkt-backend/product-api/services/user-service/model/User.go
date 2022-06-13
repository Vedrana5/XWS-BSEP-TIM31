package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	PhoneNumber string    `json:"phoneNumber" gorm:"not null"`
	FirstName   string    `json:"firstName" gorm:"not null"`
	LastName    string    `json:"lastName" gorm:"not null"`
	DateOfBirth time.Time `json:"dateOfBirth" gorm:"not null"`

	Salt           string        `json:"salt" gorm:"not null"`
	TypeOfUser     TypeOfUser    `json:"user_type" gorm:"not null"`
	TypeOfProfile  TypeOfProfile `json:"profile_type" gorm:"not null"`
	Gender         Gender        `json:"gender" gorm:";not null"`
	Biography      string        `json:"biography" gorm:"not null"`
	WorkExperience string        `json:"workExperience" gorm:"not null"`
	Education      string        `json:"education" gorm:"not null"`
	Skills         string        `json:"skills" gorm:"not null"`
	Interest       string        `json:"interest" gorm:"not null"`
	Question       string        `json:"question" gorm:"not null"`
	Answer         string        `json:"answer" gorm:"not null"`
	AnswerSalt     string        `json:"answer_salt" gorm:"not null"`
	IsConfirmed    bool          `json:"is_confirmed" gorm:"not null"`
}
