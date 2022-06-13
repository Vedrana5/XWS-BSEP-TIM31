package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"
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

func (handler *UpdateProfileHandler) UpdateUserProfileInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	if err := TokenValid(r); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("User is not logged in!")
		w.WriteHeader(http.StatusUnauthorized) // 401
		return
	}

	var EditProfileDTO dto.EditProfileDTO
	if err := json.NewDecoder(r.Body).Decode(&EditProfileDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to UserUpdateProfileInfoDTO!")
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
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("User is not authorized to update user information!")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	err := handler.UserService.UpdateUserProfileInfo(&EditProfileDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("Failed updating basic user profile information!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisteredUserHandler",
		"action":    "UPDUSPROFINF393",
		"timestamp": time.Now().String(),
	}).Info("Successfully updated user profile info!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
