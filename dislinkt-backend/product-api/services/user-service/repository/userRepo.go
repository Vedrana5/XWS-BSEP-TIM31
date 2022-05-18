package repository

import (
	"fmt"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"

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
