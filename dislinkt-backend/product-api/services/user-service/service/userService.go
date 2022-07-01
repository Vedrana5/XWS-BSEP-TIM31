package service

import (
	"github.com/google/uuid"
	"strings"
	"user/module/dto"
	"user/module/model"
	"user/module/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func (service *UserService) ChangePassword(salt string, password string, user *model.User) error {
	err := service.Repo.ChangePassword(salt, password, user)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateUserConfirmed(userId uuid.UUID, isConfirmed bool) error {
	err := service.Repo.UpdateUserConfirmed(userId, isConfirmed)
	if err != nil {
		return err
	}
	return nil
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

func (service *UserService) FindPublic(userName string) []model.User {
	users := service.Repo.FindAllUsers()
	var s []model.User
	for i := 0; i < len(users); i++ {
		if strings.Contains(users[i].Username, userName) {
			var user model.User = users[i]
			if user.TypeOfProfile == model.PUBLIC {
				s = append(s, user)
			}
		}
	}
	if s != nil {
		return s
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

func (service *UserService) FindByUserNameAndID(ID uuid.UUID, username string) *model.User {
	users := service.Repo.FindByUserNameAndID(ID, username)
	return users
}

func (service *UserService) FindByEmailAndID(ID uuid.UUID, email string) *model.User {
	users := service.Repo.FindByEmailAndID(ID, email)
	return users
}

func (service *UserService) UpdateUserProfileInfo(user *dto.EditProfileDTO) error {
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
