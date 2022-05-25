package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/gorilla/mux"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService *service.UserService
	LogInfo     *logrus.Logger
	LogError    *logrus.Logger
}

func UpdateUserConfirmed() {

}

func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var user = handler.UserService.FindByUserName(username)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FIDBYUSNAM9482",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisteredUserHandler",
		"action":    "FIDBYUSNAM9482",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindPublicByUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var userName string

	if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to UserUpdateProfileInfoDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var typeOFProfile = user.TypeOfProfile

	if typeOFProfile == model.PRIVATE {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FIDBYUSNAM9482",
			"timestamp": time.Now().String(),
		}).Error("Profile is private")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FIDBYUSNAM9482",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisteredUserHandler",
		"action":    "FIDBYUSNAM9482",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}
