package handler

import (
	"encoding/json"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"

	"github.com/sirupsen/logrus"

	"net/http"
	_ "strconv"
	"time"
)

type ConfirmationTokenHandler struct {
	ConfirmationTokenService *service.ConfirmationTokenService
	UserService              *service.UserService
	LogInfo                  *logrus.Logger
	LogError                 *logrus.Logger
}

//VERFYCONFTOK322
func (handler *ConfirmationTokenHandler) VerifyConfirmationToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var confirmationAccountDTO dto.ConfirmationAccountDTO
	err := json.NewDecoder(r.Body).Decode(&confirmationAccountDTO)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Wrong cast json to ConfirmationAccountDTO!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Confirmation token isn't valid!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Confirmation token isn't valid!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = handler.UserService.UpdateUserConfirmed(confirmationToken.UserId, true)
	if err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "ConfirmationTokenHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed updating basic user to confirmed!")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "ConfirmationTokenHandler",
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully verifed account with token!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}
