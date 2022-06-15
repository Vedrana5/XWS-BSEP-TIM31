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
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/Vedrana5/XWS-BSEP-TIM31/dislinkt-backend/product-api/services/user-service/service"
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

func (handler *UserHandler) FindByUserName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	username := vars["username"]

	var user = handler.UserService.FindByUserName(username)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	var user = handler.UserService.FindByID(uuid.MustParse(id))
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully founded user by username!")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindTokenByCode(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	confirmationToken := vars["confirmationToken"]

	var validation_token = handler.ValidationCodeService.FindByCode(uuid.MustParse(confirmationToken))
	if validation_token == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Validation token not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}

	userJson, _ := json.Marshal(validation_token)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully founded validation token by code!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Wrong cast json to string !")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}
	var user = handler.UserService.FindByUserName(userName)
	var typeOFProfile = user.TypeOfProfile

	if typeOFProfile == model.PRIVATE {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Profile is private")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User not found!")
		w.WriteHeader(http.StatusExpectationFailed)
	}
	userJson, _ := json.Marshal(user)
	w.Write(userJson)
	handler.LogInfo.WithFields(logrus.Fields{
		"status":    "success",
		"location":  "UserHandler",
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully founded user by username!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Wrong cast json to RequestDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var user = handler.UserService.FindByID(requestDTO.ID)

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error("User is not found!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Password doesn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed changing password!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *UserHandler) ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	var resetPasswordDTO dto.ResetPasswordDTO
	fmt.Print("ID USERA JE" + resetPasswordDTO.ID)
	fmt.Print("CODE JE" + resetPasswordDTO.Code)
	fmt.Print("PASSWORD JE" + resetPasswordDTO.Password)
	if err := json.NewDecoder(r.Body).Decode(&resetPasswordDTO); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Wrong cast json to RequestDTO!")
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	var user = handler.UserService.FindByID(uuid.MustParse(resetPasswordDTO.ID))

	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User is not found!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Password doesn't in valid format!")
		w.WriteHeader(http.StatusBadRequest) //400
		return
	}

	if err := handler.UserService.ChangePassword(salt, password, user); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed changing password!")
		w.WriteHeader(http.StatusExpectationFailed)
		return
	}

	var validation_code = handler.ValidationCodeService.FindByCode(uuid.MustParse(resetPasswordDTO.Code))
	if validation_code == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Validation_code is not found!")
		w.WriteHeader(http.StatusConflict) //409
		return
	}

	if err := handler.ValidationCodeService.UpdateValidationCodeUsing(validation_code.Code); err != nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed changing password!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User not found!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed creating confirmation token for user!")
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
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully sended email!")

}

func (handler *UserHandler) SendMailForPasswordlessLogin(w http.ResponseWriter, r *http.Request) {

	// Sender data.
	vars := mux.Vars(r)
	email := vars["email"]

	var user = handler.UserService.FindByEmail(email)
	if user == nil {
		handler.LogError.WithFields(logrus.Fields{
			"status":    "failure",
			"location":  "UserHandler",
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "User is not found!")
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
			"timestamp": time.Now().String(),
		}).Error(time.Now().String() + "Failed creating confirmation token for user!")
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
		"timestamp": time.Now().String(),
	}).Info(time.Now().String() + "Successfully sended email!")

}
