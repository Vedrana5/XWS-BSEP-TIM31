package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"user/module/dto"
	"user/module/model"
	"user/module/service"
	"user/module/util"

	"github.com/go-playground/validator"
	"github.com/mikespook/gorbac"
	"github.com/sirupsen/logrus"
)

type UpdateProfileHandler struct {
	Rbac                     *gorbac.RBAC
	PermissionFindAllUsers   *gorbac.Permission
	UserService              *service.UserService
	PermissionUpdateUserInfo *gorbac.Permission
	PasswordUtil             *util.PasswordUtil
	Validator                *validator.Validate
	LogInfo                  *logrus.Logger
	LogError                 *logrus.Logger
}

//UpdateUserProfileInfo
func (handler *UpdateProfileHandler) UpdateUserProfileInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("User is not logged in!")
		fmt.Println(time.Now().String() + " User is not logged in!")

		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	var EditProfileDTO dto.EditProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&EditProfileDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to UserUpdateProfileInfoDTO!")
		fmt.Println(time.Now().String() + " Wrong cast json to UserUpdateProfileInfoDTO!")

		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var loginUser = handler.UserService.FindByID(EditProfileDTO.ID)
	userRole := ""
	if loginUser.TypeOfUser == model.ADMIN {
		userRole = "role-admin"
	} else if loginUser.TypeOfUser == model.REGISTERED_USER {
		userRole = "role-registered-user"
	} else {
		userRole = "role-unregistered-user"
	}
	log.Print("ROLA JE" + userRole)
	if !handler.Rbac.IsGranted(userRole, *handler.PermissionUpdateUserInfo, nil) {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("User is not authorized to update user information!")
		fmt.Println(time.Now().String() + " User is not authorized to update user information!")

		w.WriteHeader(http.StatusForbidden)
		return
	}

	if handler.UserService.FindByEmailAndID(EditProfileDTO.ID, EditProfileDTO.Email) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered email!")
		fmt.Println(time.Now().String() + " User already exist with entered email!")

		w.WriteHeader(http.StatusExpectationFailed) //417
		return
	}

	fmt.Println("ID JE " + EditProfileDTO.ID.String())
	fmt.Println("USERNAME JE" + EditProfileDTO.Username)

	if handler.UserService.FindByUserNameAndID(EditProfileDTO.ID, EditProfileDTO.Username) != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisterHandler",
			"action":    "CreateUser",
			"timestamp": time.Now().String(),
		}).Error("User already exist with entered username!")
		fmt.Println(time.Now().String() + " User already exist with entered username!")

		w.WriteHeader(http.StatusConflict) //409
		return
	}

	err := handler.UserService.UpdateUserProfileInfo(&EditProfileDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UpdateProfileHandler",
			"action":    "UpdateUserProfileInfo",
			"timestamp": time.Now().String(),
		}).Error("Failed updating basic user profile information!")
		fmt.Println(time.Now().String() + " Failed updating basic user profile information!")

		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UpdateProfileHandler",
		"action":    "UpdateUserProfileInfo",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated user profile info!")
	fmt.Println(time.Now().String() + " Successfully updated user profile info!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
