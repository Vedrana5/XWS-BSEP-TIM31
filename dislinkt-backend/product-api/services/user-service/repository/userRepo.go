package repository

import (
	"fmt"
	"time"
	"user/module/dto"
	"user/module/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepo struct {
	Database *gorm.DB
}

func (repo *UserRepo) CreateUser(user *model.User) error {
	result := repo.Database.Create(user)
	fmt.Print(result)
	return nil
}

func (repo *UserRepo) ChangePassword(salt string, password string, user *model.User) error {

	result := repo.Database.Model(&model.User{}).Where("username = ?", user.Username)

	result.Update("password", password)
	fmt.Println(result.RowsAffected)
	result.Update("salt", salt)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating profile info")
	return nil
}

func (repo *UserRepo) FindAllUsers() []model.User {
	var users []model.User
	repo.Database.Select("*").Find(&users)
	return users
}

func (repo *UserRepo) FindByUserName(userName string) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "username = ?", userName).RowsAffected == 0 {

		return nil

	}
	return user
}

func (repo *UserRepo) FindByEmail(email string) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "email = ?", email).RowsAffected == 0 {
		return nil
	}
	return user
}

func (repo *UserRepo) FindByID(ID uuid.UUID) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "id = ?", ID).RowsAffected == 0 {
		return nil
	}
	return user
}

func (repo *UserRepo) FindByUserNameAndID(ID uuid.UUID, username string) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "username = ?", username).RowsAffected == 0 {
		return nil
	}
	if user.ID == ID {
		return nil
	}
	return user
}

func (repo *UserRepo) FindByEmailAndID(ID uuid.UUID, email string) *model.User {
	user := &model.User{}
	if repo.Database.First(&user, "email = ?", email).RowsAffected == 0 {
		return nil
	}
	if user.ID == ID {
		return nil
	}
	return user
}

func (repo *UserRepo) UpdateUserProfileInfo(user *dto.EditProfileDTO) error {
	gender := model.OTHER
	switch user.Gender {
	case "MALE":
		gender = model.MALE
	case "FEMALE":
		gender = model.FEMALE
		fmt.Println("ovdeee samm")

	}

	result := repo.Database.Model(&model.User{}).Where("username = ?", user.OldUsername)
	result.Update("username", user.Username)
	fmt.Println(result.RowsAffected)
	result.Update("phone_number", user.PhoneNumber)
	fmt.Println(result.RowsAffected)
	result.Update("first_name", user.FirstName)
	fmt.Println(result.RowsAffected)
	result.Update("last_name", user.LastName)
	fmt.Println(result.RowsAffected)
	result.Update("gender", gender)
	fmt.Println(result.RowsAffected)
	layout := "2006-01-02"
	dateOfBirth, _ := time.Parse(layout, user.DateOfBirth)
	result.Update("date_of_birth", dateOfBirth)
	fmt.Println(result.RowsAffected)
	result.Update("biography", user.Biography)
	fmt.Println(result.RowsAffected)
	result.Update("email", user.Email)
	fmt.Println(result.RowsAffected)

	result.Update("workExperience", user.WorkExperience)
	fmt.Println(result.RowsAffected)
	result.Update("education", user.Education)
	fmt.Println(result.RowsAffected)
	result.Update("skills", user.Skills)
	fmt.Println(result.RowsAffected)
	result.Update("interest", user.Interest)
	fmt.Println(result.RowsAffected)

	fmt.Println("updating profile info")
	return nil
}

func (repo *UserRepo) UpdateUserPassword(userId uuid.UUID, salt string, password string) error {
	result := repo.Database.Model(&model.User{}).Where("id = ?", userId).Update("password", password)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}

func (repo *UserRepo) UpdateUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	result := repo.Database.Model(&model.User{}).Where("id = ?", userId).Update("is_confirmed", isConfirmed)
	fmt.Println(result.RowsAffected)
	fmt.Println("updating")
	return nil
}
