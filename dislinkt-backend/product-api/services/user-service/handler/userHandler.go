package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"
	"user/module/dto"
	"user/module/model"
	"user/module/service"
	"user/module/util"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

type UserHandler struct {
	ValidationCodeService *service.ValidationCodeService
	PasswordUtil          *util.PasswordUtil
	UserService           *service.UserService
	LogInfo               *logrus.Logger
	LogError              *logrus.Logger
}

func UpdateUserConfirmed() {

}

//FindByUserName
func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var user = handler.UserService.FindByUserName(username)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindByUserName",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindByUserName",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindPublic(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var users = handler.UserService.FindPublic(username)
	if users == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindByUserName",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " Users not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(users)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindByUserName",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FindById
func (handler *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var user = handler.UserService.FindByID(uuid.MustParse(id))
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindById",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " User not found!")

		w.WriteHeader(http.StatusExpectationFailed)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindById",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FindTokenByCode
func (handler *UserHandler) FindTokenByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	confirmationToken := vars["confirmationToken"]

	var validation_token = handler.ValidationCodeService.FindByCode(uuid.MustParse(confirmationToken))
	if validation_token == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindTokenByCode",
			"timestamp": time.Now().String(),
		}).Error("Validation token not found!")
		fmt.Println(time.Now().String() + " Validation token not found!")

		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(validation_token)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindTokenByCode",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded validation token by code!")
	fmt.Println(time.Now().String() + " Successfully founded validation token by code!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

//FindPublicByUserName
func (handler *UserHandler) FindPublicByUserName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var userName string

	if err := json.NewDecoder(r.Body).Decode(&userName); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindPublicByUserName",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to string !")
		fmt.Println(time.Now().String() + " Wrong cast json to string!")

		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var typeOFProfile = user.TypeOfProfile

	if typeOFProfile == model.PRIVATE {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindPublicByUserName",
			"timestamp": time.Now().String(),
		}).Error("Profile is private")
		fmt.Println(time.Now().String() + " Profile is private!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "FindPublicByUserName",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " User not found!")

		w.WriteHeader(http.StatusExpectationFailed)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"action":    "FindPublicByUserName",
		"timestamp": time.Now().String(),
	}).Info("Successfully founded user by username!")
	fmt.Println(time.Now().String() + " Successfully founded user by username!")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

}

//ChangePassword
func (handler *UserHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var requestDTO dto.RequestDTO

	if err := json.NewDecoder(r.Body).Decode(&requestDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to RequestDTO!")
		fmt.Println(time.Now().String() + " Wrong cast json to RequestDTO!")

		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var user = handler.UserService.FindByID(requestDTO.ID)

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("User is not found!")
		fmt.Println(time.Now().String() + " User is not found!")

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
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		fmt.Println(time.Now().String() + " Password doesn't in valid format!")

		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ChangePassword",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		fmt.Println(time.Now().String() + " Failed changing password!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//ResetPassword
func (handler *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var resetPasswordDTO dto.ResetPasswordDTO
	if err := json.NewDecoder(r.Body).Decode(&resetPasswordDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("Wrong cast json to RequestDTO!")
		fmt.Println(time.Now().String() + " Wrong cast json to RequestDTO!")

		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var user = handler.UserService.FindByID(uuid.MustParse(resetPasswordDTO.ID))

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("User is not found!")
		fmt.Println(time.Now().String() + " User is not found!")

		w.WriteHeader(http.StatusConflict) //409
		return
	}

	salt := ""
	password := ""
	validPassword := handler.PasswordUtil.IsValidPassword(resetPasswordDTO.Password)

	if validPassword {
		//PASSWORD SALT
		salt, password = handler.PasswordUtil.GeneratePasswordWithSalt(resetPasswordDTO.Password)
	} else {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("Password doesn't in valid format!")
		fmt.Println(time.Now().String() + " Password doesn't in valid format!")

		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		fmt.Println(time.Now().String() + " Failed changing password!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var validation_code = handler.ValidationCodeService.FindByCode(uuid.MustParse(resetPasswordDTO.Code))
	if validation_code == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("Validation_code is not found!")
		fmt.Println(time.Now().String() + " Validation_code is not found!")

		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if err := handler.ValidationCodeService.UpdateValidationCodeUsing(validation_code.Code); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "ResetPassword",
			"timestamp": time.Now().String(),
		}).Error("Failed changing password!")
		fmt.Println(time.Now().String() + " Failed changing password!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//SendMailForResetPasword
func (handler *UserHandler) SendMailForResetPassword(w http.ResponseWriter, r *http.Request) {

	// Sender data.
	vars := mux.Vars(r)
	username := vars["username"]

	var user = handler.UserService.FindByUserName(username)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "SendMailForResetPasword",
			"timestamp": time.Now().String(),
		}).Error("User not found!")
		fmt.Println(time.Now().String() + " User not found!")

		w.WriteHeader(http.StatusExpectationFailed)
	}

	confirmationToken := model.ValidationCode{
		ID:          uuid.New(),
		Code:        uuid.New(),
		UserId:      user.ID,
		CreatedTime: time.Now(),
		ExpiredTime: time.Now().Add(time.Minute * 20),
		IsValid:     true,
		IsUsed:      false,
	}

	if err := handler.ValidationCodeService.CreateConfirmationToken(&confirmationToken); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "SendMailForResetPasword",
			"timestamp": time.Now().String(),
		}).Error("Failed creating confirmation token for user!")
		fmt.Println(time.Now().String() + " Failed creating confirmation token for user!!")

		w.WriteHeader(http.StatusExpectationFailed)
		return
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
	message := []byte("Dear " + user.FirstName + ",\n\nPlease, click on link in below to change your password on our social network!\n\nhttps://localhost:8082/confirmResetPassword/" + id.String() + "/" + confirmationToken.Code.String())

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
		"location":  "UserHandler",
		"action":    "SendMailForResetPasword",
		"timestamp": time.Now().String(),
	}).Info("Successfully sended email!")
	fmt.Println(time.Now().String() + " Successfully sended email!")

}

//SendMailForPasswordlessLogin
func (handler *UserHandler) SendMailForPasswordlessLogin(w http.ResponseWriter, r *http.Request) {

	// Sender data.
	vars := mux.Vars(r)
	email := vars["email"]

	var user = handler.UserService.FindByEmail(email)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "SendMailForPasswordlessLogin",
			"timestamp": time.Now().String(),
		}).Error("User is not found!")
		fmt.Println(time.Now().String() + " User is not found!")

		w.WriteHeader(http.StatusExpectationFailed)
	}

	confirmationToken := model.ValidationCode{
		ID:          uuid.New(),
		Code:        uuid.New(),
		UserId:      user.ID,
		CreatedTime: time.Now(),
		ExpiredTime: time.Now().Add(time.Minute * 20),
		IsValid:     true,
		IsUsed:      false,
	}

	if err := handler.ValidationCodeService.CreateConfirmationToken(&confirmationToken); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"action":    "SendMailForPasswordlessLogin",
			"timestamp": time.Now().String(),
		}).Error("Failed creating confirmation token for user!")
		fmt.Println(time.Now().String() + " Failed creating confirmation token for user!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
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
	message := []byte("Dear " + user.FirstName + ",\n\nPlease, click on link in below to change your password on our social network!\n\nhttps://localhost:8082/confirmPasswordlessLogin/" + id.String() + "/" + confirmationToken.Code.String())

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
		"location":  "UserHandler",
		"action":    "SendMailForPasswordlessLogin",
		"timestamp": time.Now().String(),
	}).Info("Successfully sended email!")
	fmt.Println(time.Now().String() + " Successfully sended email!")

}
