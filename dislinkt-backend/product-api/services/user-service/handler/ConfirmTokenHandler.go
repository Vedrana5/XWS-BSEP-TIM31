package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	_ "strconv"
	"time"
	"user/module/dto"
	"user/module/service"
)

type ConfirmationTokenHandler struct {
	ConfirmationTokenService *service.ConfirmationTokenService
	UserService              *service.UserService
	LogInfo                  *logrus.Logger
	LogError                 *logrus.Logger
}

//VerifyConfirmationToken
func (handler *ConfirmationTokenHandler) VerifyConfirmationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var confirmationAccountDTO dto.ConfirmationAccountDTO
	err := json.NewDecoder(r.Body).Decode(&confirmationAccountDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"action":    "VerifyConfirmationToken",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to ConfirmationAccountDTO!")
		fmt.Println(time.Now().String() + " Wrong cast json to ConfirmationAccountDTO!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIdUUID := confirmationAccountDTO.UserId
	tokenUUID := confirmationAccountDTO.ConfirmationToken

	var confirmationToken = handler.ConfirmationTokenService.FindByToken(tokenUUID)
	if !confirmationToken.IsValid {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"action":    "VerifyConfirmationToken",
			"timestamp": time.Now().String(),
		}).Error("Confirmation token isn't valid!")
		fmt.Println(time.Now().String() + " Confirmation token isn't valid!")

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if confirmationToken.UserId != userIdUUID || confirmationToken.ExpiredTime.Before(time.Now()) {
		err := handler.ConfirmationTokenService.UpdateConfirmationTokenValidity(confirmationToken.ConfirmationToken, false)
		if err != nil {
			return
		}
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"action":    "VerifyConfirmationToken",
			"timestamp": time.Now().String(),
		}).Error("Confirmation token isn't valid!")
		fmt.Println(time.Now().String() + " Confirmation token isn't valid!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.UpdateUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"action":    "VerifyConfirmationToken",
			"timestamp": time.Now().String(),
		}).Error("Failed updating basic user to confirmed!")
		fmt.Println(time.Now().String() + " Failed updating basic user to confirmed!")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ConfirmationTokenHandler",
		"action":    "VerifyConfirmationToken",
		"timestamp": time.Now().String(),
	}).Info("Successfully verifed account with token!")
	fmt.Println(time.Now().String() + " Successfully verifed account with token!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
