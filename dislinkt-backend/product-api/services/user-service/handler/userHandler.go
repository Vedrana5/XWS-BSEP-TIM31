package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/dto"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/model"
	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/util"
	"github.com/gorilla/mux"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	PasswordUtil *util.PasswordUtil
	UserService  *service.UserService
	LogInfo      *logrus.Logger
	LogError     *logrus.Logger
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

func (handler *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var requestDTO dto.RequestDTO

	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "UPDUSPROFINF393",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to RequestDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var user = handler.UserService.FindByID(requestDTO.ID)

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("User is null!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(requestDTO.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(requestDTO.Password)
	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "RegisteredUserHandler",
			"action":    "CRREGUS032",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) SendMailForResetPassword(w http.ResponseWriter, r *http.Request) {
	// Sender data.
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

	from := "sammilica99@gmail.com"
	password := "setmkiwpicaxhmti"

	// Receiver email address.
	to := []string{
		user.Email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	id := user.ID
	// Message.
	message := []byte("Dear " + user.FirstName + ",\n\nPlease, click on link in below to reset your password on our social network!\n\nhttp://localhost:8082/confirmResetPassword/" + id.String())

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "RegisteredUserHandler",
		"action":    "SEDCONFMAIL227",
		"timestamp": time.Now().String(),
	}).Info("Successfully sended email!")

}
