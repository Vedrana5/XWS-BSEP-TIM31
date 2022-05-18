package dto

type RegisteredUserDTO struct {
	Username       string `json:"username" validate:"required,min=2,max=30"`
	Password       string `json:"password" validate:"required,min=10,max=30"`
	Email          string `json:"email" validate:"required,email"`
	PhoneNumber    string `json:"phoneNumber" validate:"required"`
	FirstName      string `json:"firstName" validate:"required,alpha,min=2,max=20"`
	LastName       string `json:"lastName" validate:"required,alpha,min=2,max=35"`
	DateOfBirth    string `json:"dateOfBirth" validate:"required"`
	TypeOfUser     string `json:"user_type" gorm:"not null"`
	TypeOfProfile  string `json:"profile_type" gorm:"not null"`
	Gender         string `json:"gender" gorm:"unique;not null"`
	Biography      string `json:"biography" gorm:"unique;not null"`
	WorkExperience string `json:"workExperience" gorm:"unique;not null"`
	Education      string `json:"education" gorm:"unique;not null"`
	Skills         string `json:"skills" gorm:"unique;not null"`
	Interest       string `json:"interest" gorm:"unique;not null"`
}
