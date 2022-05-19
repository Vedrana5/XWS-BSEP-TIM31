package service

import (
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/repository"
	"github.com/google/uuid"
)

type UserService struct {
	Repo *repository.UserRepo
}

func (service *UserService) CreateUser(user *model.User) error {
	err := service.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) FindAllUsers() []model.User {
	users := service.Repo.FindAllUsers()
	if users != nil {
		return users
	}
	return nil
}

func (service *UserService) FindByUserName(userName string) *model.User {
	user := service.Repo.FindByUserName(userName)
	return user
}

func (service *UserService) FindByEmail(email string) *model.User {
	user := service.Repo.FindByEmail(email)
	return user
}

func (service *UserService) FindByID(ID uuid.UUID) *model.User {
	user := service.Repo.FindByID(ID)
	return user
}

func (service *UserService) UpdateUserProfileInfo(user *dto.RegisteredUserDTO) error {
	err := service.Repo.UpdateUserProfileInfo(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUserPassword(userId uuid.UUID, salt string, password string) error {
	err := service.Repo.UpdateUserPassword(userId, salt, password)
	if err != nil {
		return err
	}
	return nil
}
