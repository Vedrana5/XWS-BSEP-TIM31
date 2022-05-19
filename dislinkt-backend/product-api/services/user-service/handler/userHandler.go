package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserService *service.UserService
	LogInfo     *logrus.Logger
	LogError    *logrus.Logger
}

func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

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
	fmt.Fprintf(w, "CAO")
	var user = handler.UserService.FindByUserName(userName)
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
